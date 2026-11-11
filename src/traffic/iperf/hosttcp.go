package iperf

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"

	"kronos/models"
	"kronos/traffic"
)

// TCP tests use the host's TCP/IP: the sockets bind to the UE address, and
// the UE's TUN device routes them through the PDU session.

// Check validates the options and fills the defaults.
func (o *TCPOptions) Check() error {
	if o.Duration == 0 {
		o.Duration = models.IperfDefaultDuration
	}
	if o.Length == 0 {
		o.Length = models.IperfDefaultTCPLength
	}
	switch {
	case !o.Server.Addr().Is4() || o.Server.Port() == 0:
		return fmt.Errorf("the iperf server %s must be an IPv4 address and a port", o.Server)
	case !o.Local.Is4():
		return fmt.Errorf("the local address %s is not IPv4", o.Local)
	case o.Rate > models.IperfMaxRate:
		return fmt.Errorf("the iperf rate is at most %s", traffic.FormatRate(models.IperfMaxRate))
	case o.Duration < time.Second || o.Duration > models.IperfMaxDuration || o.Duration%time.Second != 0:
		return fmt.Errorf("the iperf time is whole seconds from 1s to %s (%s)", models.IperfMaxDuration, o.Duration)
	case o.Length < 1 || o.Length > models.IperfMaxTCPLength:
		return fmt.Errorf("the iperf TCP len is 1 to %d octets (%d)", models.IperfMaxTCPLength, o.Length)
	}
	return nil
}

// String renders the summary of the test.
func (r TCPResult) String() string {
	dir := "uplink (UE → DN)"
	if r.Reverse {
		dir = "downlink (DN → UE)"
	}
	secs := r.Seconds
	if secs <= 0 {
		secs = 1
	}
	s := fmt.Sprintf("TCP %s %.1f s: sent %s (%s), received %s (%s)", dir, r.Seconds,
		traffic.FormatRate(float64(r.SentBytes)*8/secs), traffic.FormatBytes(r.SentBytes),
		traffic.FormatRate(float64(r.RecvBytes)*8/secs), traffic.FormatBytes(r.RecvBytes))
	if r.Retransmits >= 0 {
		s += fmt.Sprintf(", retransmits %d", r.Retransmits)
	}
	return s
}

// RunTCP makes one TCP test from o.Local against the server.
func RunTCP(ctx context.Context, o TCPOptions, report Reporter) (TCPResult, error) {
	if err := o.Check(); err != nil {
		return TCPResult{}, err
	}
	// The control connection and the setup: cookie, then parameters.
	d := net.Dialer{LocalAddr: &net.TCPAddr{IP: o.Local.AsSlice()}, Timeout: connectTimeout}
	cc, err := d.DialContext(ctx, "tcp4", o.Server.String())
	if err != nil {
		return TCPResult{}, fmt.Errorf("cannot connect to %s (is iperf3 -s running?): %v", o.Server, err)
	}
	defer cc.Close()
	k := hostControl(cc)
	cookie := makeCookie()
	if _, err := cc.Write(cookie); err != nil {
		return TCPResult{}, err
	}
	if err := k.expect(paramExchange, stepTimeout); err != nil {
		return TCPResult{}, err
	}
	p := params{TCP: true, Time: int(o.Duration / time.Second), Parallel: 1, Reverse: o.Reverse,
		Len: o.Length, Bandwidth: o.Rate, PacingTimer: pacingTimer, ClientVersion: clientVersion}
	if err := k.writeJSON(p); err != nil {
		return TCPResult{}, err
	}
	if err := k.expect(createStreams, stepTimeout); err != nil {
		return TCPResult{}, err
	}
	// The data connection, which starts with the same cookie.
	dconn, err := d.DialContext(ctx, "tcp4", o.Server.String())
	if err != nil {
		return TCPResult{}, fmt.Errorf("cannot open the data connection: %v", err)
	}
	dc := dconn.(*net.TCPConn)
	defer dc.Close()
	if _, err := dc.Write(cookie); err != nil {
		return TCPResult{}, err
	}
	if err := k.expect(testStart, stepTimeout); err != nil {
		return TCPResult{}, err
	}
	if err := k.expect(testRunning, stepTimeout); err != nil {
		return TCPResult{}, err
	}

	// The test: the data moves on a goroutine until the time is up. The
	// server says nothing on the control connection before TEST_END, so a
	// state that comes during the test ends it.
	var moved atomic.Uint64
	start := time.Now()
	deadline := start.Add(o.Duration)
	stop := make(chan struct{})
	done := make(chan error, 1)
	go func() { done <- pumpTCP(dc, o, start, deadline, &moved, stop) }()
	ctl := make(chan ctlRead, 1)
	k.deadline(time.Time{})
	go func() {
		var b [1]byte
		n, err := cc.Read(b[:])
		ctl <- ctlRead{n, b[0], err}
	}()
	endData := func() {
		close(stop)
		dc.SetDeadline(time.Now())
		<-done
	}
	endCtl := func() (ctlRead, bool) { // stops the reader; true if a state came
		k.deadline(time.Now())
		r := <-ctl
		return r, r.n == 1
	}
	verb := "sent"
	if o.Reverse {
		verb = "received"
	}
	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	second, last := 1, uint64(0)
	interval := func() {
		cur := moved.Load()
		report(false, fmt.Sprintf("%d–%d s: %s %s", second-1, second, verb, traffic.FormatRate(float64(cur-last)*8)))
		last = cur
		second++
	}
	var dataErr error
	var elapsed float64
loop:
	for {
		select {
		case <-ctx.Done():
			endData()
			endCtl()
			k.writeState(clientTerminate)
			return TCPResult{}, ErrAborted
		case <-tick.C:
			interval()
		case dataErr = <-done:
			elapsed = time.Since(start).Seconds()
			break loop
		case r := <-ctl:
			endData()
			if r.n == 0 {
				return TCPResult{}, fmt.Errorf("the control connection ended during the test: %v", r.err)
			}
			if int8(r.b) == serverError {
				return TCPResult{}, errors.New(k.serverErrorText())
			}
			return TCPResult{}, fmt.Errorf("the server sent %s during the test", stateName(int8(r.b)))
		}
	}
	if second <= int(o.Duration/time.Second) {
		interval() // the last second, when the data ended before its tick
	}
	if r, came := endCtl(); came {
		return TCPResult{}, fmt.Errorf("the server sent %s during the test", stateName(int8(r.b)))
	}
	if dataErr != nil {
		return TCPResult{}, fmt.Errorf("the data connection failed: %v", dataErr)
	}
	retrans := int64(-1)
	if !o.Reverse {
		retrans = retransmits(dc)
	} else {
		// What the server sent before it read TEST_END is drained, not counted.
		go io.Copy(io.Discard, dc)
	}

	// The end: exchange the results.
	if err := k.writeState(testEnd); err != nil {
		return TCPResult{}, err
	}
	if err := k.expect(exchangeResults, resultsTimeout); err != nil {
		return TCPResult{}, err
	}
	mine := streamResult{ID: 1, Bytes: float64(moved.Load()), Retransmits: float64(retrans), EndTime: elapsed}
	has := -1 // the receiving end has no retransmits to tell
	if !o.Reverse {
		has = 0
		if retrans >= 0 {
			has = 1
		}
	}
	if err := k.writeJSON(results{SenderHasRetransmits: has, Streams: []streamResult{mine}}); err != nil {
		return TCPResult{}, err
	}
	var theirs results
	if err := k.readJSON(&theirs, resultsTimeout); err != nil {
		return TCPResult{}, fmt.Errorf("cannot read the results of the server: %v", err)
	}
	if len(theirs.Streams) == 0 {
		return TCPResult{}, errors.New("the results of the server have no stream")
	}
	if err := k.expect(displayResults, resultsTimeout); err != nil {
		return TCPResult{}, err
	}
	k.writeState(iperfDone)

	// Combine our side with the server's.
	server := theirs.Streams[0]
	r := TCPResult{Reverse: o.Reverse, Seconds: elapsed, Retransmits: retrans}
	if o.Reverse {
		r.SentBytes, r.RecvBytes, r.Retransmits = uint64(server.Bytes), moved.Load(), -1
		if theirs.SenderHasRetransmits == 1 {
			r.Retransmits = int64(server.Retransmits)
		}
	} else {
		r.SentBytes, r.RecvBytes = moved.Load(), uint64(server.Bytes)
	}
	return r, nil
}

// pumpTCP writes (uplink) or reads (downlink) the data until the deadline
// or stop, counting the octets in moved.
func pumpTCP(c *net.TCPConn, o TCPOptions, start, deadline time.Time, moved *atomic.Uint64, stop <-chan struct{}) error {
	if o.Reverse {
		c.SetReadDeadline(deadline)
		buf := make([]byte, max(o.Length, tcpReadBufLen))
		for {
			n, err := c.Read(buf)
			moved.Add(uint64(n))
			if err != nil {
				return endOfData(err)
			}
		}
	}
	c.SetWriteDeadline(deadline)
	buf := make([]byte, o.Length)
	for i := range buf {
		buf[i] = byte(i)
	}
	for {
		if o.Rate > 0 { // pace the writes
			due := start.Add(time.Duration(float64(moved.Load()+uint64(o.Length)) * 8 / float64(o.Rate) * float64(time.Second)))
			if wait := time.Until(due); wait > 0 {
				if due.After(deadline) {
					return nil
				}
				t := time.NewTimer(wait)
				select {
				case <-t.C:
				case <-stop:
					t.Stop()
					return nil
				}
			}
		}
		n, err := c.Write(buf)
		moved.Add(uint64(n))
		if err != nil {
			return endOfData(err)
		}
	}
}

// endOfData: the deadline (the end of the test, or stop) and the peer's
// close end the data without an error.
func endOfData(err error) error {
	if errors.Is(err, os.ErrDeadlineExceeded) || errors.Is(err, io.EOF) {
		return nil
	}
	return err
}

// retransmits reads the segments the kernel sent again (TCP_INFO), or -1.
func retransmits(c *net.TCPConn) int64 {
	raw, err := c.SyscallConn()
	if err != nil {
		return -1
	}
	got := int64(-1)
	raw.Control(func(fd uintptr) {
		var info syscall.TCPInfo
		size := uint32(syscall.SizeofTCPInfo)
		_, _, e := syscall.Syscall6(syscall.SYS_GETSOCKOPT, fd, syscall.IPPROTO_TCP, syscall.TCP_INFO,
			uintptr(unsafe.Pointer(&info)), uintptr(unsafe.Pointer(&size)), 0)
		if e == 0 {
			got = int64(info.Total_retrans)
		}
	})
	return got
}
