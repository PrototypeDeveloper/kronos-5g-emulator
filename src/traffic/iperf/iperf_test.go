package iperf

import (
	"bytes"
	"context"
	"io"
	"math/rand/v2"
	"net"
	"net/netip"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"kronos/models"
	"kronos/traffic"
	"kronos/traffic/tcp"
)

var (
	ueAddr = netip.MustParseAddr("10.60.0.7")
	dnAddr = netip.MustParseAddr("10.10.10.20")
)

// wire connects a UE stack and a DN stack; loss drops that share of the
// packets both ways. Packets are delivered on another goroutine, as the
// GTP-U reader does.
func wire(t *testing.T, loss float64) (ue, dn *traffic.Stack, ueLink, dnLink traffic.Link) {
	ue, dn = traffic.NewStack(), traffic.NewStack()
	var mu sync.Mutex
	r := rand.New(rand.NewPCG(3, 4))
	pipe := func(to *traffic.Stack) traffic.Link {
		ch := make(chan []byte, 8192)
		done := make(chan struct{})
		go func() {
			for {
				select {
				case p := <-ch:
					to.Deliver(p)
				case <-done:
					return
				}
			}
		}()
		t.Cleanup(func() { close(done) })
		return traffic.LinkFunc(func(pkt []byte) error {
			mu.Lock()
			drop := r.Float64() < loss
			mu.Unlock()
			if !drop {
				select {
				case ch <- append([]byte(nil), pkt...):
				default: // a full queue drops, as a real link
				}
			}
			return nil
		})
	}
	return ue, dn, pipe(dn), pipe(ue)
}

func run(t *testing.T, loss float64, reverse bool) Result {
	t.Helper()
	ue, dn, ueLink, dnLink := wire(t, loss)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go Serve(ctx, dn, dnLink, dnAddr, 13531)
	var lines []string
	var mu sync.Mutex
	r, err := Run(ctx, ue, ueLink, ueAddr, Options{Server: netip.AddrPortFrom(dnAddr, 13531), Rate: 8_000_000, Duration: 2 * time.Second, Reverse: reverse, Length: 1000},
		func(warn bool, text string) {
			mu.Lock()
			lines = append(lines, text)
			mu.Unlock()
		})
	if err != nil {
		t.Fatalf("iperf: %v (lines %q)", err, lines)
	}
	mu.Lock()
	defer mu.Unlock()
	if len(lines) != 2 || !strings.HasPrefix(lines[0], "0–1 s: ") {
		t.Errorf("interval lines %q", lines)
	}
	t.Log(r)
	return r
}

func TestUplink(t *testing.T) {
	r := run(t, 0, false)
	// 8 Mbit/s for 2 s of 1000-octet datagrams: about 2000 datagrams.
	if r.Reverse || r.SentPackets < 1900 || r.SentPackets > 2100 || r.RecvBytes != r.SentBytes || r.Lost != 0 || r.Total != r.SentPackets {
		t.Fatalf("result %+v", r)
	}
}

func TestDownlink(t *testing.T) {
	r := run(t, 0, true)
	if !r.Reverse || r.SentPackets < 1900 || r.RecvBytes != r.SentBytes || r.Lost != 0 {
		t.Fatalf("result %+v", r)
	}
}

func TestLoss(t *testing.T) {
	for _, reverse := range []bool{false, true} {
		r := run(t, 0.02, reverse)
		loss := float64(r.Lost) / float64(r.Total)
		if loss < 0.005 || loss > 0.05 || r.RecvBytes >= r.SentBytes {
			t.Fatalf("reverse %v: loss %.3f, result %+v", reverse, loss, r)
		}
	}
}

func TestOptions(t *testing.T) {
	for _, o := range []Options{
		{Server: netip.AddrPortFrom(dnAddr, 0)},
		{Server: netip.AddrPortFrom(dnAddr, 1), Duration: 1500 * time.Millisecond},
		{Server: netip.AddrPortFrom(dnAddr, 1), Length: 8},
		{Server: netip.AddrPortFrom(dnAddr, 1), Rate: models.IperfMaxRate + 1},
	} {
		if err := o.Check(); err == nil {
			t.Errorf("%+v accepted", o)
		}
	}
	o := Options{Server: netip.AddrPortFrom(dnAddr, 5201)}
	if err := o.Check(); err != nil || o.Rate != models.IperfDefaultRate || o.Duration != models.IperfDefaultDuration || o.Length != models.IperfDefaultLength {
		t.Fatalf("defaults %+v %v", o, err)
	}
}

// TestRealIperf3 runs the client against the iperf3 of this host, through a
// relay between the DN stack and kernel sockets (no privileges needed).
func TestRealIperf3(t *testing.T) {
	bin, err := exec.LookPath("iperf3")
	if err != nil {
		t.Skip("iperf3 is not installed")
	}
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	port := l.Addr().(*net.TCPAddr).Port
	l.Close()
	var out bytes.Buffer
	srv := exec.Command(bin, "-s", "-B", "127.0.0.1", "-p", strconv.Itoa(port))
	srv.Stdout, srv.Stderr, srv.Dir = &out, &out, t.TempDir()
	if err := srv.Start(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		srv.Process.Kill()
		srv.Wait()
		if t.Failed() {
			t.Logf("iperf3 -s:\n%s", out.String())
		}
	}()
	for i := 0; i < 50; i++ { // wait for it to listen
		if c, err := net.Dial("tcp", "127.0.0.1:"+strconv.Itoa(port)); err == nil {
			c.Close()
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	// The probe above was a test of its own for iperf3; give it a moment to reset.
	time.Sleep(300 * time.Millisecond)

	ue, dn, ueLink, dnLink := wire(t, 0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	relay(t, ctx, dn, dnLink, uint16(port))
	for _, reverse := range []bool{false, true} {
		r, err := Run(ctx, ue, ueLink, ueAddr, Options{Server: netip.AddrPortFrom(dnAddr, uint16(port)), Rate: 10_000_000, Duration: 2 * time.Second, Reverse: reverse},
			func(bool, string) {})
		if err != nil {
			t.Fatalf("reverse %v: %v", reverse, err)
		}
		t.Log(r)
		if r.SentPackets < 1500 || r.RecvBytes == 0 || r.Lost > r.Total/20 {
			t.Fatalf("reverse %v: result %+v", reverse, r)
		}
	}
}

// relay forwards the iperf3 control connection and the UDP stream between
// dnAddr:port of the DN stack and 127.0.0.1:port of the kernel.
func relay(t *testing.T, ctx context.Context, dn *traffic.Stack, dnLink traffic.Link, port uint16) {
	l, err := tcp.Listen(dn, dnLink, port)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(l.Close)
	target := "127.0.0.1:" + strconv.Itoa(int(port))
	go func() {
		for {
			c, err := l.Accept(ctx)
			if err != nil {
				return
			}
			k, err := net.Dial("tcp", target)
			if err != nil {
				c.Abort()
				continue
			}
			go func() { io.Copy(k, c); k.(*net.TCPConn).CloseWrite() }()
			go func() { io.Copy(c, k); c.Close() }()
		}
	}()
	var mu sync.Mutex
	var udp *net.UDPConn
	var peer netip.AddrPort
	dn.Bind(models.ProtoUDP, port, func(ip models.IPv4) {
		u, ok := traffic.ParseUDP(ip)
		if !ok {
			return
		}
		mu.Lock()
		defer mu.Unlock()
		if udp == nil || u.Src != peer {
			if udp != nil {
				udp.Close()
			}
			a, _ := net.ResolveUDPAddr("udp", target)
			c, err := net.DialUDP("udp", nil, a)
			if err != nil {
				return
			}
			udp, peer = c, u.Src
			go func(c *net.UDPConn, to netip.AddrPort) {
				buf := make([]byte, 65536)
				for {
					n, err := c.Read(buf)
					if err != nil {
						return
					}
					dnLink.Send(traffic.BuildUDP(netip.AddrPortFrom(dnAddr, port), to, dn.IPID(), buf[:n]))
				}
			}(c, u.Src)
		}
		udp.Write(u.Payload)
	})
	t.Cleanup(func() {
		mu.Lock()
		if udp != nil {
			udp.Close()
		}
		mu.Unlock()
	})
}
