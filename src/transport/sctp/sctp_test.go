package sctp

import (
	"bytes"
	"errors"
	"io"
	"kronos/models"
	"net/netip"
	"os"
	"testing"
	"time"
)

func listenLoopback(t *testing.T) *Listener {
	t.Helper()
	l, err := Listen(netip.MustParseAddrPort("127.0.0.1:0"), Options{})
	if err != nil {
		t.Skipf("SCTP is not available here: %v", err)
	}
	t.Cleanup(func() { l.Close() })
	return l
}

func TestDialSendRecv(t *testing.T) {
	l := listenLoopback(t)
	accepted := make(chan *Conn, 1)
	go func() {
		c, err := l.Accept()
		if err != nil {
			t.Error(err)
			close(accepted)
			return
		}
		accepted <- c
	}()
	c, err := Dial(netip.MustParseAddrPort("127.0.0.1:0"), l.Addr(), Options{OutStreams: 4, InStreams: 4})
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	s := <-accepted
	if s == nil {
		t.FailNow()
	}
	defer s.Close()
	if out, in := c.Streams(); out != 4 || in != 4 {
		t.Fatalf("streams out=%d in=%d", out, in)
	}
	if c.RemoteAddr() != l.Addr() || !c.LocalAddr().IsValid() {
		t.Fatalf("addresses %v → %v", c.LocalAddr(), c.RemoteAddr())
	}

	big := bytes.Repeat([]byte{0xab}, 100_000) // larger than one receive buffer
	for _, m := range []struct {
		b      []byte
		stream uint16
	}{{[]byte{0x00, 0x15, 0x00}, 0}, {[]byte("ue"), 3}, {big, 1}} {
		if err := c.Send(m.b, m.stream, models.PPIDNGAP); err != nil {
			t.Fatal(err)
		}
		got, info, err := s.Recv()
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(got, m.b) || info.Stream != m.stream || info.PPID != models.PPIDNGAP {
			t.Fatalf("got %d bytes stream %d ppid %d", len(got), info.Stream, info.PPID)
		}
	}

	// The peer closing the association is reported as io.EOF.
	c.Close()
	if _, _, err := s.Recv(); !errors.Is(err, io.EOF) {
		t.Fatalf("after peer close: %v", err)
	}
}

func TestCloseUnblocksRecv(t *testing.T) {
	l := listenLoopback(t)
	go func() {
		if c, err := l.Accept(); err == nil {
			time.Sleep(time.Second)
			c.Close()
		}
	}()
	c, err := Dial(netip.AddrPort{}, l.Addr(), Options{})
	if err != nil {
		t.Fatal(err)
	}
	done := make(chan error, 1)
	go func() { _, _, err := c.Recv(); done <- err }()
	time.Sleep(50 * time.Millisecond)
	c.Close()
	select {
	case err := <-done:
		if !errors.Is(err, os.ErrClosed) {
			t.Fatalf("Recv after Close: %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Recv still blocked after Close")
	}
}

func TestDialRefused(t *testing.T) {
	l := listenLoopback(t)
	addr := l.Addr()
	l.Close()
	if _, err := Dial(netip.AddrPort{}, addr, Options{Timeout: 2 * time.Second}); err == nil {
		t.Fatal("dial to a closed port succeeded")
	}
}
