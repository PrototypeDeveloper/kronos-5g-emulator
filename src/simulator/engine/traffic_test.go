package engine_test

import (
	"context"
	"net"
	"net/netip"
	"strings"
	"testing"

	libgtp "libgtp"
	"libgtp/gtpMessage"
	"libgtp/gtpType"

	"kronos/models"
	"kronos/traffic"
	"kronos/traffic/iperf"
)

// dnBehindUPF emulates what is behind the UPF: the inner packets of uplink
// G-PDUs go to a DN stack (pings are answered), and the DN's packets come
// back to the gNB in G-PDUs with the downlink TEID.
func dnBehindUPF(t *testing.T, upf *net.UDPConn, gnb *net.UDPAddr, dlTEID uint32) (*traffic.Stack, traffic.Link) {
	st := traffic.NewStack()
	link := traffic.LinkFunc(func(pkt []byte) error {
		g, err := gtpMessage.NewGPDUWithPDUSessionContainer(dlTEID, pkt, gtpType.PDUSessionInformationDL, 1)
		if err != nil {
			return err
		}
		b, err := g.Encode()
		if err != nil {
			return err
		}
		_, err = upf.WriteToUDP(b, gnb)
		return err
	})
	go func() {
		buf := make([]byte, 65536)
		for {
			n, err := upf.Read(buf)
			if err != nil {
				return
			}
			msg, err := libgtp.Decode(buf[:n])
			if err != nil {
				continue
			}
			g, ok := msg.(*gtpMessage.GPDU)
			if !ok || g.Header.TEID != ulTEID {
				continue
			}
			pkt := append([]byte(nil), g.TPDU...)
			if ip, ok := traffic.ParseIPv4(pkt); ok {
				if e, ok := traffic.ParseEcho(ip); ok && e.Type == 8 {
					link.Send(traffic.BuildEcho(ip.Dst, ip.Src, 1, 0, e.ID, e.Seq, e.Data))
					continue
				}
			}
			st.Deliver(pkt)
		}
	}()
	return st, link
}

// TestTrafficThroughN3 runs ping and iperf through the engine's GTP-U tunnel
// to an emulated DN: ping with its options and summary, iperf uplink and
// downlink against an iperf3 server, and stop.
func TestTrafficThroughN3(t *testing.T) {
	upf, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		t.Fatal(err)
	}
	defer upf.Close()
	gtpuPort := freeUDPPort(t)
	hs := startN3(t, true, 1, sessionScenario, gtpuPort, upf.LocalAddr().(*net.UDPAddr).Port)
	hs.linkup(t)
	amfSec, ran := register(t, hs)
	dlTEID := establish(t, hs, amfSec, ran)
	dn, dnLink := dnBehindUPF(t, upf, &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: gtpuPort}, dlTEID)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go iperf.Serve(ctx, dn, dnLink, netip.MustParseAddr("10.10.10.20"), 13531)
	const imsi = "target(208930000000001)."
	exec := func(line string) {
		t.Helper()
		if _, err := hs.eng.Exec(imsi + line); err != nil {
			t.Fatalf("%s: %v", line, err)
		}
	}
	text := func(what, substr string) string {
		t.Helper()
		return hs.wait(t, what, func(e models.Event) bool { return strings.Contains(e.Text, substr) }).Text
	}

	exec("ping(10.10.10.20, 3, 0.2s, 100)")
	if s := text("ping summary", "ping 10.10.10.20: 3 transmitted, 3 received, 0% loss"); !strings.Contains(s, "rtt min/avg/max/mdev = ") {
		t.Fatalf("summary %q", s)
	}

	exec("iperf(10.10.10.20, 13531, 4M, 1s)")
	text("iperf start", "iperf 10.10.10.20:13531: uplink (UE → DN) 4.0 Mbit/s × 1s, 1400 octets, starting (UE 10.60.0.7)")
	if s := text("iperf uplink", "iperf 10.10.10.20:13531: uplink (UE → DN) "); !strings.Contains(s, ", lost 0/") {
		t.Fatalf("uplink %q", s)
	}
	exec("iperf(10.10.10.20, 13531, 4M, 1s, dl, 1000)")
	text("iperf start", "iperf 10.10.10.20:13531: downlink (DN → UE) 4.0 Mbit/s × 1s, 1000 octets, starting")
	if s := text("iperf downlink", "iperf 10.10.10.20:13531: downlink (DN → UE) "); !strings.Contains(s, ", lost 0/") {
		t.Fatalf("downlink %q", s)
	}

	// Runs that go on until stop.
	exec("ping(10.10.10.20, 0, 50ms)")
	text("ping progress", "ping 10.10.10.20: 1 s: ")
	exec("stop(ping)")
	text("stopped ping", " loss, time ")
	exec("iperf(10.10.10.20, 13531, 1M, 30s)")
	text("iperf running", "iperf 10.10.10.20:13531: 0–1 s: sent ")
	exec("stop")
	text("stopped iperf", "iperf 10.10.10.20:13531: stopped")

	// iperftcp needs the UE's TUN device (ue.tun).
	if _, err := hs.eng.Exec(imsi + "iperftcp(10.10.10.20, 13531)"); err == nil || !strings.Contains(err.Error(), "set tun: true under ue:") {
		t.Fatalf("iperftcp without ue.tun: %v", err)
	}

	// Without a session there is nothing to run on.
	if _, err := hs.eng.Exec("linkdown"); err != nil {
		t.Fatal(err)
	}
	text("reset", "back to the initial state")
	for _, line := range []string{"ping(10.10.10.20)", "iperf(10.10.10.20)", "iperftcp(10.10.10.20)"} {
		if _, err := hs.eng.Exec(imsi + line); err == nil || !strings.Contains(err.Error(), "no usable PDU session") {
			t.Fatalf("%s without a session: %v", line, err)
		}
	}
}
