package iperf

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/netip"
	"sync/atomic"
	"time"

	"kronos/models"
	"kronos/traffic"
	"kronos/traffic/tcp"
)

// Serve answers iperf3 UDP tests on addr:port of a stack, one test at a
// time, until ctx ends. It is the far end for tests (an emulated DN): only
// what Run uses is implemented.
func Serve(ctx context.Context, st *traffic.Stack, link traffic.Link, addr netip.Addr, port uint16) error {
	l, err := tcp.Listen(st, link, port)
	if err != nil {
		return err
	}
	defer l.Close()
	for {
		c, err := l.Accept(ctx)
		if err != nil {
			return err
		}
		serveOne(ctx, st, link, netip.AddrPortFrom(addr, port), c)
	}
}

// serveOne runs one test on the accepted control connection c.
func serveOne(ctx context.Context, st *traffic.Stack, link traffic.Link, self netip.AddrPort, c *tcp.Conn) error {
	// The setup: cookie and parameters; only UDP tests are supported.
	k := smallControl(c)
	defer c.Close()
	c.SetReadDeadline(time.Now().Add(stepTimeout))
	if _, err := io.ReadFull(c, make([]byte, cookieLen)); err != nil {
		return err
	}
	if err := k.writeState(paramExchange); err != nil {
		return err
	}
	var p params
	if err := k.readJSON(&p, stepTimeout); err != nil {
		return err
	}
	if !p.UDP || p.Len < udpHeader || p.Time < 1 {
		k.writeState(serverError)
		c.Write([]byte{0, 0, 0, 1, 0, 0, 0, 0}) // i_errno 1, errno 0
		return fmt.Errorf("unsupported parameters %+v", p)
	}

	// The data stream: the client's first datagram gives its address.
	recv := &receiver{}
	var client atomic.Pointer[netip.AddrPort]
	ready := make(chan struct{})
	if _, err := st.Bind(models.ProtoUDP, self.Port(), func(ip models.IPv4) {
		u, ok := traffic.ParseUDP(ip)
		if !ok {
			return
		}
		if client.Load() == nil {
			src := u.Src
			if client.CompareAndSwap(nil, &src) {
				close(ready)
			}
			return
		}
		if u.Src == *client.Load() && !p.Reverse {
			recv.take(u.Payload, time.Now())
		}
	}); err != nil {
		return err
	}
	defer st.Unbind(models.ProtoUDP, self.Port())
	if err := k.writeState(createStreams); err != nil {
		return err
	}
	select {
	case <-ready:
	case <-time.After(stepTimeout):
		return errors.New("no UDP connect")
	}
	// Answer the UDP connect and start the test.
	peer := *client.Load()
	link.Send(traffic.BuildUDP(self, peer, st.IPID(), udpConnectReply))
	if err := k.writeState(testStart); err != nil {
		return err
	}
	if err := k.writeState(testRunning); err != nil {
		return err
	}

	// The test: send in the background (reverse) until the client sends TEST_END.
	snd := newSender(p.Bandwidth, p.Len, func(payload []byte) error {
		return link.Send(dataPacket(self, peer, st.IPID(), payload))
	})
	stop := make(chan struct{})
	sent := make(chan struct{})
	start := time.Now()
	go func() {
		defer close(sent)
		if !p.Reverse {
			return
		}
		snd.start = start
		tick := time.NewTicker(tickInterval)
		defer tick.Stop()
		for {
			select {
			case <-stop:
				return
			case <-ctx.Done():
				return
			case now := <-tick.C:
				snd.pump(now)
			}
		}
	}()
	s, err := k.readState(time.Duration(p.Time+testEndSlack) * time.Second)
	close(stop)
	<-sent
	if err != nil {
		return err
	}
	if s != testEnd {
		return fmt.Errorf("got %s", stateName(s))
	}
	// The end: exchange the results and wait for IPERF_DONE.
	elapsed := time.Since(start).Seconds()
	if err := k.writeState(exchangeResults); err != nil {
		return err
	}
	var theirs results
	if err := k.readJSON(&theirs, stepTimeout); err != nil {
		return err
	}
	mine := streamResult{ID: 1, Retransmits: -1, EndTime: elapsed}
	if p.Reverse {
		mine.Bytes, mine.Packets = float64(snd.bytes), float64(snd.packets)
	} else {
		rs := recv.snapshot()
		mine.Bytes, mine.Jitter, mine.Errors, mine.Packets = float64(rs.bytes), rs.jitter, float64(rs.lost), float64(rs.count)
	}
	if err := k.writeJSON(results{Streams: []streamResult{mine}}); err != nil {
		return err
	}
	if err := k.writeState(displayResults); err != nil {
		return err
	}
	return k.expect(iperfDone, stepTimeout)
}
