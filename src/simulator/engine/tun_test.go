package engine_test

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"net/netip"
	"strings"
	"sync"
	"testing"
	"time"

	libnas "libnas"
	"libnas/nasMessage"
	"libngap/ngapMessage"
	"libngap/ngapType"

	"kronos/models"
	"kronos/simulator/engine"
	"kronos/traffic"
)

// fakeTUN stands for a host TUN device: host holds what the host sends,
// toHost gets what the engine writes.
type fakeTUN struct {
	name   string
	host   chan []byte
	toHost chan []byte
	closed chan struct{}
	once   sync.Once
}

func (f *fakeTUN) Name() string { return f.name }

func (f *fakeTUN) Read(b []byte) (int, error) {
	select {
	case p := <-f.host:
		return copy(b, p), nil
	case <-f.closed:
		return 0, io.EOF
	}
}

func (f *fakeTUN) Write(b []byte) (int, error) {
	select {
	case f.toHost <- append([]byte(nil), b...):
	default:
	}
	return len(b), nil
}

func (f *fakeTUN) Close() error {
	f.once.Do(func() { close(f.closed) })
	return nil
}

func (f *fakeTUN) isClosed() bool {
	select {
	case <-f.closed:
		return true
	default:
		return false
	}
}

// fakeTUNs opens fake devices and remembers them.
type fakeTUNs struct {
	mu     sync.Mutex
	opened []*fakeTUN
	addrs  []netip.Addr
}

func useFakeTUN(hs *harness) *fakeTUNs {
	ft := &fakeTUNs{}
	engine.UseTUN(hs.eng, func(index int, ueIP netip.Addr) (engine.TUNDevice, error) {
		f := &fakeTUN{name: fmt.Sprintf("kronos-ue%d", index+1), host: make(chan []byte, 16), toHost: make(chan []byte, 16), closed: make(chan struct{})}
		ft.mu.Lock()
		ft.opened = append(ft.opened, f)
		ft.addrs = append(ft.addrs, ueIP)
		ft.mu.Unlock()
		return f, nil
	})
	return ft
}

func (ft *fakeTUNs) count() int {
	ft.mu.Lock()
	defer ft.mu.Unlock()
	return len(ft.opened)
}

func (ft *fakeTUNs) last() *fakeTUN {
	ft.mu.Lock()
	defer ft.mu.Unlock()
	return ft.opened[len(ft.opened)-1]
}

// echoRequest builds an IPv4 ICMP echo request.
func echoRequest(src, dst netip.Addr, id, seq uint16) []byte {
	icmp := make([]byte, 8+16)
	icmp[0] = 8
	binary.BigEndian.PutUint16(icmp[4:], id)
	binary.BigEndian.PutUint16(icmp[6:], seq)
	copy(icmp[8:], "kronos tun test!")
	binary.BigEndian.PutUint16(icmp[2:], inetChecksum(icmp))
	ip := make([]byte, 20, 20+len(icmp))
	ip[0], ip[8], ip[9] = 0x45, 64, 1
	binary.BigEndian.PutUint16(ip[2:], uint16(20+len(icmp)))
	s, d := src.As4(), dst.As4()
	copy(ip[12:], s[:])
	copy(ip[16:], d[:])
	binary.BigEndian.PutUint16(ip[10:], inetChecksum(ip))
	return append(ip, icmp...)
}

// TestTUN checks the data path of ue.tun: the host's packets from the UE
// address go uplink through the session, downlink packets that no built-in
// endpoint takes come back to the host, and the built-in ping keeps its own.
func TestTUN(t *testing.T) {
	upf, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		t.Fatal(err)
	}
	defer upf.Close()
	gtpuPort := freeUDPPort(t)
	hs := startN3(t, true, 1, sessionScenario, gtpuPort, upf.LocalAddr().(*net.UDPAddr).Port)
	ft := useFakeTUN(hs)
	hs.linkup(t)
	amfSec, ran := register(t, hs)
	dlTEID := establish(t, hs, amfSec, ran)
	dnBehindUPF(t, upf, &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: gtpuPort}, dlTEID)
	hs.wait(t, "TUN up", func(e models.Event) bool {
		return strings.Contains(e.Text, "TUN kronos-ue1 up: 10.60.0.7, MTU 1456 (the host's packets from 10.60.0.7 go through PDU session 1")
	})
	if ft.count() != 1 || ft.addrs[0] != netip.MustParseAddr("10.60.0.7") {
		t.Fatalf("opened %d %v", ft.count(), ft.addrs)
	}
	dev := ft.last()
	ue, dn := netip.MustParseAddr("10.60.0.7"), netip.MustParseAddr("10.10.10.20")

	// A packet from another address is dropped; the host's ping goes out and its reply comes back.
	dev.host <- echoRequest(netip.MustParseAddr("10.60.0.8"), dn, 0x4242, 1)
	dev.host <- echoRequest(ue, dn, 0x4242, 2)
	select {
	case p := <-dev.toHost:
		ip, ok := traffic.ParseIPv4(p)
		e, eok := traffic.ParseEcho(ip)
		if !ok || !eok || ip.Src != dn || ip.Dst != ue || e.Type != 0 || e.ID != 0x4242 || e.Seq != 2 || string(e.Data) != "kronos tun test!" {
			t.Fatalf("to the host: %x", p)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("no echo reply to the host")
	}

	// The built-in ping takes its replies; the host does not see them.
	if _, err := hs.eng.Exec("target(208930000000001).ping(10.10.10.20, 2, 0.1s)"); err != nil {
		t.Fatal(err)
	}
	hs.wait(t, "ping summary", func(e models.Event) bool { return strings.Contains(e.Text, "2 transmitted, 2 received") })
	select {
	case p := <-dev.toHost:
		t.Fatalf("the host got %x", p)
	default:
	}

	// linkdown removes the device.
	if _, err := hs.eng.Exec("linkdown"); err != nil {
		t.Fatal(err)
	}
	hs.wait(t, "TUN removed", func(e models.Event) bool {
		return e.Text == "TUN kronos-ue1 removed (1 uplink packet(s) dropped: no tunnel or not from the UE)"
	})
	if !dev.isClosed() {
		t.Fatal("the device is open after linkdown")
	}
}

// TestDeregistrationAccept: when the AMF accepts a de-registration, the UE
// releases its PDU sessions locally; their tunnels and the TUN device go.
func TestDeregistrationAccept(t *testing.T) {
	hs := startN3(t, true, 1, sessionScenario, freeUDPPort(t), freeUDPPort(t))
	ft := useFakeTUN(hs)
	hs.linkup(t)
	amfSec, ran := register(t, hs)
	establish(t, hs, amfSec, ran)
	hs.wait(t, "TUN up", func(e models.Event) bool { return strings.HasPrefix(e.Text, "TUN kronos-ue1 up") })

	nas, err := libnas.Encode(&nasMessage.DeregistrationAcceptUE{},
		&libnas.EncodeContext{Security: amfSec, SecurityHeaderType: libnas.SecurityHeaderIntegrityProtectedAndCiphered, Direction: libnas.Downlink})
	if err != nil {
		t.Fatal(err)
	}
	hs.amf.send(t, &ngapMessage.DownlinkNASTransport{AMFUENGAPID: ngapType.AMFUENGAPID{Value: 100}, RANUENGAPID: ngapType.RANUENGAPID{Value: ran}, NASPDU: ngapType.NASPDU{Value: nas}})
	hs.wait(t, "local release", func(e models.Event) bool { return e.Text == "PDU session(s) 1 released locally (deregistered)" })
	hs.wait(t, "TUN removed", func(e models.Event) bool { return strings.HasPrefix(e.Text, "TUN kronos-ue1 removed") })
	if _, err := hs.eng.Exec("target(208930000000001).ping(10.10.10.20)"); err == nil || !strings.Contains(err.Error(), "no usable PDU session") {
		t.Fatalf("ping after the de-registration: %v", err)
	}
	if st, _ := hs.eng.UE("208930000000001"); len(st.Sessions) != 0 || st.RM != models.RMDeregistered || !ft.last().isClosed() {
		t.Fatalf("after the de-registration: %+v", st)
	}
}
