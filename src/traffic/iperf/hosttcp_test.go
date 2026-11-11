package iperf

import (
	"bytes"
	"context"
	"errors"
	"io"
	"kronos/models"
	"net"
	"net/netip"
	"os/exec"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

var loopback = netip.MustParseAddr("127.0.0.1")

// serveTCPHost answers one iperf3 TCP test on l, as iperf3 -s does: only
// what RunTCP uses. It returns the octets the data connection carried.
func serveTCPHost(l net.Listener) (uint64, error) {
	cc, err := l.Accept()
	if err != nil {
		return 0, err
	}
	defer cc.Close()
	k := hostControl(cc)
	cookie := make([]byte, cookieLen)
	cc.SetReadDeadline(time.Now().Add(5 * time.Second))
	if _, err := io.ReadFull(cc, cookie); err != nil {
		return 0, err
	}
	k.writeState(paramExchange)
	var p params
	if err := k.readJSON(&p, 5*time.Second); err != nil {
		return 0, err
	}
	if !p.TCP || p.Len < 1 || p.Time < 1 || p.Parallel != 1 {
		return 0, errors.New("bad parameters")
	}
	k.writeState(createStreams)
	dc, err := l.Accept()
	if err != nil {
		return 0, err
	}
	defer dc.Close()
	got := make([]byte, cookieLen)
	dc.SetReadDeadline(time.Now().Add(5 * time.Second))
	if _, err := io.ReadFull(dc, got); err != nil || !bytes.Equal(got, cookie) {
		return 0, errors.New("bad data cookie")
	}
	dc.SetReadDeadline(time.Time{})
	k.writeState(testStart)
	k.writeState(testRunning)
	var moved atomic.Uint64
	go func() { // send (reverse) or receive until TEST_END closes the test
		buf := make([]byte, p.Len)
		for {
			var n int
			var err error
			if p.Reverse {
				n, err = dc.Write(buf)
			} else {
				n, err = dc.Read(buf)
			}
			moved.Add(uint64(n))
			if err != nil {
				return
			}
		}
	}()
	if err := k.expect(testEnd, time.Duration(p.Time+5)*time.Second); err != nil {
		return 0, err
	}
	dc.SetDeadline(time.Now())
	k.writeState(exchangeResults)
	var theirs results
	if err := k.readJSON(&theirs, 5*time.Second); err != nil {
		return 0, err
	}
	has := -1
	if p.Reverse {
		has = 1
	}
	k.writeJSON(results{SenderHasRetransmits: has, Streams: []streamResult{{ID: 1, Bytes: float64(moved.Load()), Retransmits: 7}}})
	k.writeState(displayResults)
	if err := k.expect(iperfDone, 5*time.Second); err != nil {
		return 0, err
	}
	return moved.Load(), nil
}

func listenLoopback(t *testing.T) net.Listener {
	l, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { l.Close() })
	return l
}

func TestTCP(t *testing.T) {
	for _, reverse := range []bool{false, true} {
		l := listenLoopback(t)
		served := make(chan uint64, 1)
		go func() {
			n, err := serveTCPHost(l)
			if err != nil {
				t.Errorf("server: %v", err)
			}
			served <- n
		}()
		var lines []string
		o := TCPOptions{Server: l.Addr().(*net.TCPAddr).AddrPort(), Local: loopback, Rate: 80_000_000, Duration: 2 * time.Second, Reverse: reverse, Length: 16 << 10}
		r, err := RunTCP(context.Background(), o, func(_ bool, s string) { lines = append(lines, s) })
		if err != nil {
			t.Fatalf("reverse %v: %v", reverse, err)
		}
		n := <-served
		t.Log(r, lines)
		if len(lines) != 2 || !strings.HasPrefix(lines[0], "0–1 s: ") || !strings.HasPrefix(lines[1], "1–2 s: ") {
			t.Fatalf("progress %q", lines)
		}
		if reverse {
			// The server is not paced here, so it sends what loopback takes.
			if r.SentBytes != n || r.RecvBytes == 0 || r.RecvBytes > r.SentBytes || r.Retransmits != 7 || !strings.Contains(r.String(), "TCP downlink (DN → UE) ") {
				t.Fatalf("downlink %+v (server %d)", r, n)
			}
			continue
		}
		// Paced at 80 Mbit/s for 2 s: about 20 MB.
		if r.SentBytes < 18_000_000 || r.SentBytes > 21_000_000 || r.RecvBytes == 0 || r.RecvBytes != n || r.Retransmits < 0 ||
			!strings.Contains(r.String(), "TCP uplink (UE → DN) ") || !strings.Contains(r.String(), ", retransmits ") {
			t.Fatalf("uplink %+v (server %d)", r, n)
		}
	}
}

func TestTCPStopAndErrors(t *testing.T) {
	l := listenLoopback(t)
	go serveTCPHost(l)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(1200 * time.Millisecond)
		cancel()
	}()
	start := time.Now()
	_, err := RunTCP(ctx, TCPOptions{Server: l.Addr().(*net.TCPAddr).AddrPort(), Local: loopback, Duration: 30 * time.Second}, func(bool, string) {})
	if !errors.Is(err, ErrAborted) || time.Since(start) > 5*time.Second {
		t.Fatalf("stop: %v after %s", err, time.Since(start))
	}

	for _, o := range []TCPOptions{
		{Server: netip.MustParseAddrPort("127.0.0.1:0"), Local: loopback},
		{Server: netip.MustParseAddrPort("127.0.0.1:1"), Local: netip.MustParseAddr("::1")},
		{Server: netip.MustParseAddrPort("127.0.0.1:1"), Local: loopback, Length: models.IperfMaxTCPLength + 1},
		{Server: netip.MustParseAddrPort("127.0.0.1:1"), Local: loopback, Duration: 1500 * time.Millisecond},
	} {
		if err := o.Check(); err == nil {
			t.Errorf("%+v accepted", o)
		}
	}
	o := TCPOptions{Server: netip.MustParseAddrPort("127.0.0.1:5201"), Local: loopback}
	if err := o.Check(); err != nil || o.Rate != 0 || o.Duration != models.IperfDefaultDuration || o.Length != models.IperfDefaultTCPLength {
		t.Fatalf("defaults %+v %v", o, err)
	}
}

// TestTCPRealIperf3 runs a TCP test against the iperf3 of this host on
// loopback, both ways.
func TestTCPRealIperf3(t *testing.T) {
	bin, err := exec.LookPath("iperf3")
	if err != nil {
		t.Skip("iperf3 is not installed")
	}
	l := listenLoopback(t)
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
	server := netip.AddrPortFrom(loopback, uint16(port))
	// iperf3 closes its listening socket after each test and opens it again
	// a moment later: a connection in between is refused or reset.
	run := func(reverse bool) TCPResult {
		t.Helper()
		for i := 0; ; i++ {
			r, err := RunTCP(context.Background(), TCPOptions{Server: server, Local: loopback, Rate: 50_000_000, Duration: 2 * time.Second, Reverse: reverse}, func(bool, string) {})
			if err == nil {
				t.Log(r)
				return r
			}
			if i == 50 || !strings.Contains(err.Error(), "cannot connect") && !strings.Contains(err.Error(), "connection reset") {
				t.Fatalf("reverse %v: %v", reverse, err)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
	if r := run(false); r.SentBytes < 11_000_000 || r.SentBytes > 13_500_000 || r.RecvBytes == 0 || r.RecvBytes > r.SentBytes {
		t.Fatalf("uplink %+v", r)
	}
	time.Sleep(500 * time.Millisecond)
	if r := run(true); r.SentBytes < 11_000_000 || r.SentBytes > 13_500_000 || r.RecvBytes == 0 || r.RecvBytes > r.SentBytes || r.Retransmits < 0 {
		t.Fatalf("downlink %+v", r)
	}
}
