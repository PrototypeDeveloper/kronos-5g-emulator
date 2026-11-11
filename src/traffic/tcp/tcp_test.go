package tcp

import (
	"context"
	"io"
	"math/rand/v2"
	"net/netip"
	"sync"
	"testing"
	"time"

	"kronos/models"
	"kronos/traffic"
)

var (
	ueAddr = netip.MustParseAddr("10.60.0.7")
	dnAddr = netip.MustParseAddr("10.10.10.20")
)

// wire connects two stacks; loss drops that share of the packets (both ways).
func wire(t *testing.T, loss float64) (ue, dn *traffic.Stack, ueLink, dnLink traffic.Link) {
	ue, dn = traffic.NewStack(), traffic.NewStack()
	var mu sync.Mutex
	r := rand.New(rand.NewPCG(1, 2))
	drop := func() bool {
		mu.Lock()
		defer mu.Unlock()
		return r.Float64() < loss
	}
	// Deliver on another goroutine, as the GTP-U reader does.
	pipe := func(to *traffic.Stack) traffic.Link {
		ch := make(chan []byte, 1024)
		go func() {
			for p := range ch {
				to.Deliver(p)
			}
		}()
		t.Cleanup(func() { close(ch) })
		return traffic.LinkFunc(func(pkt []byte) error {
			if !drop() {
				ch <- append([]byte(nil), pkt...)
			}
			return nil
		})
	}
	return ue, dn, pipe(dn), pipe(ue)
}

func exchange(t *testing.T, loss float64) {
	ue, dn, ueLink, dnLink := wire(t, loss)
	l, err := Listen(dn, dnLink, 5201)
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	big := make([]byte, 50_000) // several windows of segments
	for i := range big {
		big[i] = byte(i * 7)
	}
	served := make(chan error, 1)
	go func() {
		c, err := l.Accept(ctx)
		if err != nil {
			served <- err
			return
		}
		got := make([]byte, len(big))
		if _, err := io.ReadFull(c, got); err != nil {
			served <- err
			return
		}
		for i := range got {
			if got[i] != big[i] {
				served <- io.ErrUnexpectedEOF
				return
			}
		}
		if _, err := c.Write([]byte("thanks")); err != nil {
			served <- err
			return
		}
		// Read until the client closes, then close too.
		if _, err := c.Read(make([]byte, 1)); err != io.EOF {
			served <- err
			return
		}
		served <- c.Close()
	}()

	c, err := Dial(ctx, ue, ueLink, ueAddr, netip.AddrPortFrom(dnAddr, 5201))
	if err != nil {
		t.Fatal(err)
	}
	if c.LocalAddr().Port() < 49152 || c.RemoteAddr().Port() != 5201 {
		t.Fatalf("addresses %s → %s", c.LocalAddr(), c.RemoteAddr())
	}
	if _, err := c.Write(big); err != nil {
		t.Fatal(err)
	}
	reply := make([]byte, 6)
	c.SetReadDeadline(time.Now().Add(10 * time.Second))
	if _, err := io.ReadFull(c, reply); err != nil || string(reply) != "thanks" {
		t.Fatalf("reply %q %v", reply, err)
	}
	if err := c.Close(); err != nil {
		t.Fatal(err)
	}
	if err := <-served; err != nil {
		t.Fatalf("server: %v", err)
	}
	if _, err := c.Write([]byte("x")); err == nil {
		t.Fatal("write after close")
	}
}

func TestExchange(t *testing.T)         { exchange(t, 0) }
func TestExchangeWithLoss(t *testing.T) { exchange(t, 0.1) }

func TestDialRefusedAndTimeout(t *testing.T) {
	ue, _, ueLink, _ := wire(t, 0)
	// Nobody listens: the DN stack drops the SYN, the dial times out with the context.
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	if _, err := Dial(ctx, ue, ueLink, ueAddr, netip.AddrPortFrom(dnAddr, 9)); err == nil {
		t.Fatal("dial without a listener succeeded")
	}
}

func TestReadDeadline(t *testing.T) {
	ue, dn, ueLink, dnLink := wire(t, 0)
	l, _ := Listen(dn, dnLink, 7)
	defer l.Close()
	ctx := context.Background()
	go l.Accept(ctx)
	c, err := Dial(ctx, ue, ueLink, ueAddr, netip.AddrPortFrom(dnAddr, 7))
	if err != nil {
		t.Fatal(err)
	}
	c.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
	if _, err := c.Read(make([]byte, 1)); err == nil {
		t.Fatal("read did not time out")
	}
	c.Abort()
}

// TestDialPastTimeWait: the peer still holds an old connection of the same
// ports and answers the SYN with an ACK of it. The client resets that and
// its next SYN opens the connection.
func TestDialPastTimeWait(t *testing.T) {
	ue, dn, ueLink, dnLink := wire(t, 0)
	const old = 0x12345678
	var mu sync.Mutex
	syns, rst := 0, false
	self := netip.AddrPortFrom(dnAddr, 5201)
	dn.Bind(models.ProtoTCP, 5201, func(ip models.IPv4) {
		_, s, ok := parseSegment(ip)
		if !ok {
			return
		}
		mu.Lock()
		defer mu.Unlock()
		switch {
		case s.flags&rstFlag != 0:
			if s.seq == old && s.flags&ackFlag == 0 {
				rst = true // (the reset of the Abort below is another one)
			}
		case s.flags&synFlag != 0:
			syns++
			if !rst { // TIME_WAIT: the ACK of the old connection
				dnLink.Send(buildSegment(1, self, s.src, 777, old, ackFlag, nil))
			} else {
				dnLink.Send(buildSegment(1, self, s.src, 1000, s.seq+1, synFlag|ackFlag, nil))
			}
		}
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	c, err := Dial(ctx, ue, ueLink, ueAddr, self)
	if err != nil {
		t.Fatal(err)
	}
	c.Abort()
	mu.Lock()
	defer mu.Unlock()
	if !rst || syns != 2 {
		t.Fatalf("reset %v, %d SYNs", rst, syns)
	}
}
