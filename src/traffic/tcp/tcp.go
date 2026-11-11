// Package tcp is a small TCP for the emulated UEs: enough for
// the control connection of iperf3. It opens connections (actively and
// passively), carries an in-order byte stream with retransmission and closes
// them. It is not meant for bulk data: a few segments are in flight, and
// out-of-order segments are dropped (the peer sends them again).
package tcp

import (
	"context"
	"encoding/binary"
	"errors"
	"io"
	"math/rand/v2"
	"net/netip"
	"os"
	"time"

	"kronos/models"
	"kronos/traffic"
)

// Errors of a connection.
var (
	ErrTimeout = errors.New("TCP: no answer from the peer")
	ErrReset   = errors.New("TCP: the peer reset the connection")
	ErrClosed  = errors.New("TCP: the connection is closed")
)

// state is the phase of a connection; closing is tracked by the FIN fields.
type state uint8

// seqLT and seqLE compare sequence numbers modulo 2^32.
func seqLT(a, b uint32) bool { return int32(a-b) < 0 }
func seqLE(a, b uint32) bool { return int32(a-b) <= 0 }

// parseSegment reads a TCP segment (checksum verified); the payload is copied.
func parseSegment(ip models.IPv4) (dstPort uint16, s segment, ok bool) {
	p := ip.Payload
	if ip.Proto != models.ProtoTCP || ip.Frag || len(p) < 20 {
		return 0, s, false
	}
	off := int(p[12]>>4) * 4
	if off < 20 || off > len(p) || traffic.TransportChecksum(ip.Src, ip.Dst, models.ProtoTCP, p) != 0 {
		return 0, s, false
	}
	s.src = netip.AddrPortFrom(ip.Src, binary.BigEndian.Uint16(p[0:]))
	s.seq = binary.BigEndian.Uint32(p[4:])
	s.ack = binary.BigEndian.Uint32(p[8:])
	s.flags = p[13]
	s.window = binary.BigEndian.Uint16(p[14:])
	s.payload = append([]byte(nil), p[off:]...)
	return binary.BigEndian.Uint16(p[2:]), s, true
}

// buildSegment returns an IPv4 packet with one TCP segment. A SYN announces MSS.
func buildSegment(ipID uint16, src, dst netip.AddrPort, seq, ack uint32, flags uint8, payload []byte) []byte {
	hl := 20
	if flags&synFlag != 0 {
		hl = 24
	}
	b := traffic.AppendIPv4(make([]byte, 0, 20+hl+len(payload)), src.Addr(), dst.Addr(), models.ProtoTCP, ipID, true, hl+len(payload))
	t := make([]byte, hl)
	binary.BigEndian.PutUint16(t[0:], src.Port())
	binary.BigEndian.PutUint16(t[2:], dst.Port())
	binary.BigEndian.PutUint32(t[4:], seq)
	binary.BigEndian.PutUint32(t[8:], ack)
	t[12] = byte(hl/4) << 4
	t[13] = flags
	binary.BigEndian.PutUint16(t[14:], recvWindow)
	if hl == 24 {
		t[20], t[21] = 2, 4 // MSS option
		binary.BigEndian.PutUint16(t[22:], MSS)
	}
	b = append(b, t...)
	b = append(b, payload...)
	binary.BigEndian.PutUint16(b[36:], traffic.TransportChecksum(src.Addr(), dst.Addr(), models.ProtoTCP, b[20:]))
	return b
}

// newConn returns an unopened connection with a random initial sequence number.
func newConn(st *traffic.Stack, link traffic.Link) *Conn {
	iss := rand.Uint32()
	return &Conn{st: st, link: link, in: make(chan segment, inQueueLen), done: make(chan struct{}),
		changed: make(chan struct{}), iss: iss, sndUna: iss, sndNxt: iss + 1, peerWnd: MSS, rto: rtoInitial}
}

// LocalAddr and RemoteAddr are the endpoints of the connection.
func (c *Conn) LocalAddr() netip.AddrPort  { return c.local }
func (c *Conn) RemoteAddr() netip.AddrPort { return c.remote }

// notify wakes the goroutines waiting for a change (mu held).
func (c *Conn) notify() {
	close(c.changed)
	c.changed = make(chan struct{})
}

// send sends one segment that acknowledges what has arrived (mu held).
func (c *Conn) send(seq uint32, flags uint8, payload []byte) {
	c.link.Send(buildSegment(c.st.IPID(), c.local, c.remote, seq, c.rcvNxt, flags, payload))
}

// sendAck sends a bare ACK (mu held).
func (c *Conn) sendAck() { c.send(c.sndNxt, ackFlag, nil) }

// push hands a received segment to the connection's goroutine.
func (c *Conn) push(s segment) {
	select {
	case c.in <- s:
	default: // a full queue drops the segment; the peer sends it again
	}
}

// Dial opens a connection from src to dst (an ephemeral local port).
func Dial(ctx context.Context, st *traffic.Stack, link traffic.Link, src netip.Addr, dst netip.AddrPort) (*Conn, error) {
	c := newConn(st, link)
	c.remote = dst
	port, err := st.Bind(models.ProtoTCP, 0, func(ip models.IPv4) {
		if _, s, ok := parseSegment(ip); ok && s.src == dst {
			c.push(s)
		}
	})
	if err != nil {
		return nil, err
	}
	c.local = netip.AddrPortFrom(src, port)
	c.release = func() { st.Unbind(models.ProtoTCP, port) }
	// Send the SYN and wait until the handshake completes or fails.
	c.mu.Lock()
	c.st8 = synSent
	c.send(c.iss, synFlag, nil)
	c.lastSend = time.Now()
	c.mu.Unlock()
	go c.loop()
	for {
		c.mu.Lock()
		if c.st8 == established {
			c.mu.Unlock()
			return c, nil
		}
		if c.err != nil {
			err := c.err
			c.mu.Unlock()
			return nil, err
		}
		ch := c.changed
		c.mu.Unlock()
		select {
		case <-ch:
		case <-c.done:
		case <-ctx.Done():
			c.Abort()
			return nil, ctx.Err()
		}
	}
}

// loop runs the connection: received segments and the retransmission timer.
func (c *Conn) loop() {
	tick := time.NewTicker(timerTick)
	defer tick.Stop()
	for {
		select {
		case s := <-c.in:
			c.mu.Lock()
			c.handle(s)
		case now := <-tick.C:
			c.mu.Lock()
			c.timer(now)
		}
		end := c.err != nil || (c.finAcked && (c.rcvFIN || time.Since(c.finAckedAt) > finWait))
		if end && c.err == nil {
			c.err = ErrClosed // later reads and writes fail rather than wait
		}
		c.mu.Unlock()
		if end {
			c.release()
			close(c.done)
			c.mu.Lock()
			c.notify()
			c.mu.Unlock()
			return
		}
	}
}

// fail ends the connection with err, keeping the first error (mu held).
func (c *Conn) fail(err error) {
	if c.err == nil {
		c.err = err
	}
	c.notify()
}

// handle processes one received segment (mu held).
func (c *Conn) handle(s segment) {
	if s.flags&rstFlag != 0 {
		if c.st8 != synSent || (s.flags&ackFlag != 0 && s.ack == c.iss+1) {
			c.fail(ErrReset)
		}
		return
	}
	switch c.st8 {
	case synSent:
		if s.flags&ackFlag != 0 && s.ack != c.iss+1 {
			// It acknowledges something else, such as a connection of the
			// same ports that the peer still holds (TIME_WAIT). A reset
			// clears that, and the SYN sent again opens this one (RFC 9293
			// 3.10.7.3).
			c.link.Send(buildSegment(c.st.IPID(), c.local, c.remote, s.ack, 0, rstFlag, nil))
			return
		}
		if s.flags&(synFlag|ackFlag) == synFlag|ackFlag {
			c.rcvNxt = s.seq + 1
			c.sndUna, c.peerWnd = s.ack, uint32(s.window)
			c.toEstablished()
			c.sendAck()
		}
		return
	case synRcvd:
		if s.flags&synFlag != 0 { // our SYN-ACK was lost
			c.send(c.iss, synFlag|ackFlag, nil)
			return
		}
		if s.flags&ackFlag == 0 || s.ack != c.iss+1 {
			return
		}
		c.sndUna, c.peerWnd = s.ack, uint32(s.window)
		c.toEstablished()
		if c.onEstablished != nil {
			c.onEstablished()
		}
	default:
		if s.flags&synFlag != 0 { // the peer did not get our ACK of its SYN
			c.sendAck()
			return
		}
	}
	if s.flags&ackFlag != 0 {
		c.acked(s.ack, s.window)
	}
	// Data and FIN are taken only in order.
	if len(s.payload) > 0 || s.flags&finFlag != 0 {
		if s.seq == c.rcvNxt && !c.rcvFIN {
			c.rcvBuf = append(c.rcvBuf, s.payload...)
			c.rcvNxt += uint32(len(s.payload))
			if s.flags&finFlag != 0 {
				c.rcvNxt++
				c.rcvFIN = true
			}
			c.notify()
		}
		c.sendAck() // in order or not: a duplicate ACK asks for what is missing
	}
}

// toEstablished completes the opening and sends any queued data (mu held).
func (c *Conn) toEstablished() {
	c.st8 = established
	c.retries, c.rto, c.lastSend = 0, rtoInitial, time.Time{}
	c.notify()
	c.output()
}

// acked takes an acknowledgement (mu held).
func (c *Conn) acked(ack uint32, wnd uint16) {
	c.peerWnd = uint32(wnd)
	if !seqLT(c.sndUna, ack) || !seqLE(ack, c.sndNxt) {
		c.output() // the window may have opened
		return
	}
	n := ack - c.sndUna
	c.sndBuf = c.sndBuf[min(int(n), len(c.sndBuf)):]
	c.sndUna = ack
	if c.finSent && ack == c.sndNxt && !c.finAcked {
		c.finAcked, c.finAckedAt = true, time.Now()
	}
	c.retries, c.rto = 0, rtoInitial
	c.lastSend = time.Time{}
	if c.sndUna != c.sndNxt {
		c.lastSend = time.Now()
	}
	c.notify()
	c.output()
}

// output sends what the window allows: data, then the FIN (mu held).
func (c *Conn) output() {
	if c.st8 != established {
		return
	}
	wnd := min(c.peerWnd, sendWindow)
	inFlight := min(c.sndNxt-c.sndUna, uint32(len(c.sndBuf)))
	for inFlight < uint32(len(c.sndBuf)) && inFlight < wnd {
		n := min(uint32(len(c.sndBuf))-inFlight, MSS, wnd-inFlight)
		if c.sndUna == c.sndNxt {
			c.lastSend = time.Now()
		}
		c.send(c.sndNxt, ackFlag|pshFlag, c.sndBuf[inFlight:inFlight+n])
		c.sndNxt += n
		inFlight += n
	}
	if c.finQueued && !c.finSent && inFlight == uint32(len(c.sndBuf)) {
		if c.sndUna == c.sndNxt {
			c.lastSend = time.Now()
		}
		c.send(c.sndNxt, finFlag|ackFlag, nil)
		c.sndNxt++
		c.finSent = true
	}
}

// timer retransmits what is overdue: go back to the oldest unacknowledged
// octet and send again (mu held).
func (c *Conn) timer(now time.Time) {
	outstanding := c.st8 != established || c.sndUna != c.sndNxt
	if !outstanding || c.lastSend.IsZero() || now.Sub(c.lastSend) < c.rto {
		return
	}
	c.retries++
	if c.retries > maxRetries {
		c.fail(ErrTimeout)
		return
	}
	c.rto = min(2*c.rto, rtoMax)
	c.lastSend = now
	switch c.st8 {
	case synSent:
		c.send(c.iss, synFlag, nil)
	case synRcvd:
		c.send(c.iss, synFlag|ackFlag, nil)
	default:
		c.sndNxt, c.finSent = c.sndUna, false
		c.output()
		c.lastSend = now
	}
}

// Write sends b and returns once the peer has acknowledged all of it.
func (c *Conn) Write(b []byte) (int, error) {
	c.mu.Lock()
	switch {
	case c.err != nil:
		err := c.err
		c.mu.Unlock()
		return 0, err
	case c.finQueued:
		c.mu.Unlock()
		return 0, ErrClosed
	}
	c.sndBuf = append(c.sndBuf, b...)
	target := c.sndUna + uint32(len(c.sndBuf))
	c.output()
	for {
		if seqLE(target, c.sndUna) {
			c.mu.Unlock()
			return len(b), nil
		}
		if c.err != nil {
			err := c.err
			c.mu.Unlock()
			return 0, err
		}
		ch := c.changed
		c.mu.Unlock()
		select {
		case <-ch:
		case <-c.done:
		}
		c.mu.Lock()
	}
}

// Read reads what has arrived in order; io.EOF after the peer's FIN.
func (c *Conn) Read(b []byte) (int, error) {
	c.mu.Lock()
	for {
		if len(c.rcvBuf) > 0 {
			n := copy(b, c.rcvBuf)
			c.rcvBuf = c.rcvBuf[n:]
			c.mu.Unlock()
			return n, nil
		}
		switch {
		case c.rcvFIN:
			c.mu.Unlock()
			return 0, io.EOF
		case c.err != nil:
			err := c.err
			c.mu.Unlock()
			return 0, err
		}
		// Nothing yet: wait for a change, the end or the deadline.
		ch, deadline := c.changed, c.readTimeout
		c.mu.Unlock()
		var t *time.Timer
		var timeout <-chan time.Time
		if !deadline.IsZero() {
			d := time.Until(deadline)
			if d <= 0 {
				return 0, os.ErrDeadlineExceeded
			}
			t = time.NewTimer(d)
			timeout = t.C
		}
		select {
		case <-ch:
		case <-c.done:
		case <-timeout:
			return 0, os.ErrDeadlineExceeded
		}
		if t != nil {
			t.Stop()
		}
		c.mu.Lock()
	}
}

// Pending reports whether a Read would return at once: data has arrived,
// or the connection ended (the peer closed it, or it failed).
func (c *Conn) Pending() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.rcvBuf) > 0 || c.rcvFIN || c.err != nil
}

// SetReadDeadline bounds the following reads (zero: no limit).
func (c *Conn) SetReadDeadline(t time.Time) {
	c.mu.Lock()
	c.readTimeout = t
	c.notify()
	c.mu.Unlock()
}

// Close sends a FIN after the data and waits (briefly) for the closing
// handshake; a peer that does not answer gets a reset.
func (c *Conn) Close() error {
	c.mu.Lock()
	if !c.finQueued && c.err == nil {
		c.finQueued = true
		c.output()
	}
	c.mu.Unlock()
	select {
	case <-c.done:
	case <-time.After(rtoMax + finWait):
		c.Abort()
	}
	return nil
}

// Abort resets the connection at once.
func (c *Conn) Abort() {
	c.mu.Lock()
	if c.err == nil {
		c.send(c.sndNxt, rstFlag|ackFlag, nil)
		c.fail(ErrClosed)
	}
	c.mu.Unlock()
	select {
	case c.in <- segment{}: // wake the loop so it ends
	default:
	}
	<-c.done
}

// Listen accepts connections to port on the stack's UE address(es).
func Listen(st *traffic.Stack, link traffic.Link, port uint16) (*Listener, error) {
	l := &Listener{st: st, link: link, port: port, accept: make(chan *Conn, acceptQueueLen), conns: map[netip.AddrPort]*Conn{}, closed: make(chan struct{})}
	if _, err := st.Bind(models.ProtoTCP, port, l.deliver); err != nil {
		return nil, err
	}
	return l, nil
}

// deliver passes a segment to its connection; a SYN from a new peer opens one.
func (l *Listener) deliver(ip models.IPv4) {
	_, s, ok := parseSegment(ip)
	if !ok {
		return
	}
	l.mu.Lock()
	c := l.conns[s.src]
	if c == nil {
		if s.flags&(synFlag|ackFlag|rstFlag) != synFlag {
			l.mu.Unlock()
			return
		}
		// A new peer: answer its SYN with a SYN-ACK.
		c = newConn(l.st, l.link)
		c.local, c.remote = netip.AddrPortFrom(ip.Dst, l.port), s.src
		c.st8, c.rcvNxt, c.peerWnd = synRcvd, s.seq+1, uint32(s.window)
		c.onEstablished = func() {
			select {
			case l.accept <- c:
			default: // nobody accepts: the connection stays unused
			}
		}
		src := s.src
		c.release = func() {
			l.mu.Lock()
			delete(l.conns, src)
			l.mu.Unlock()
		}
		l.conns[s.src] = c
		c.mu.Lock()
		c.send(c.iss, synFlag|ackFlag, nil)
		c.lastSend = time.Now()
		c.mu.Unlock()
		l.mu.Unlock()
		go c.loop()
		return
	}
	l.mu.Unlock()
	c.push(s)
}

// Accept waits for the next established connection.
func (l *Listener) Accept(ctx context.Context) (*Conn, error) {
	select {
	case c := <-l.accept:
		return c, nil
	case <-l.closed:
		return nil, ErrClosed
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// Close stops accepting. The connections of the port stop receiving too:
// close them first.
func (l *Listener) Close() {
	l.once.Do(func() {
		l.st.Unbind(models.ProtoTCP, l.port)
		close(l.closed)
	})
}
