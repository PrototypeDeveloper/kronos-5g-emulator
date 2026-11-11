package iperf

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/netip"
	"sync/atomic"
	"time"

	"kronos/models"
	"kronos/traffic"
	"kronos/traffic/tcp"
)

// String renders the summary of the test.
func (r Result) String() string {
	dir := "uplink (UE → DN)"
	if r.Reverse {
		dir = "downlink (DN → UE)"
	}
	secs := r.Seconds
	if secs <= 0 {
		secs = 1
	}
	loss := 0.0
	if r.Total > 0 {
		loss = 100 * float64(r.Lost) / float64(r.Total)
	}
	s := fmt.Sprintf("%s %.1f s: sent %s (%s), received %s (%s), lost %d/%d (%.3g%%), jitter %.3f ms", dir, r.Seconds,
		traffic.FormatRate(float64(r.SentBytes)*8/secs), traffic.FormatBytes(r.SentBytes),
		traffic.FormatRate(float64(r.RecvBytes)*8/secs), traffic.FormatBytes(r.RecvBytes),
		r.Lost, r.Total, loss, r.Jitter)
	if r.OutOfOrder > 0 {
		s += fmt.Sprintf(", out of order %d", r.OutOfOrder)
	}
	return s
}

// ErrAborted is returned when ctx ends during the test (stop).
var ErrAborted = errors.New("stopped")

// Reporter receives the progress lines; warn marks a problem.
type Reporter func(warn bool, text string)

// Check validates the options and fills the defaults.
func (o *Options) Check() error {
	if o.Rate == 0 {
		o.Rate = models.IperfDefaultRate
	}
	if o.Duration == 0 {
		o.Duration = models.IperfDefaultDuration
	}
	if o.Length == 0 {
		o.Length = models.IperfDefaultLength
	}
	switch {
	case !o.Server.Addr().Is4() || o.Server.Port() == 0:
		return fmt.Errorf("the iperf server %s must be an IPv4 address and a port", o.Server)
	case o.Rate > models.IperfMaxRate:
		return fmt.Errorf("the iperf rate is at most %s", traffic.FormatRate(models.IperfMaxRate))
	case o.Duration < time.Second || o.Duration > models.IperfMaxDuration || o.Duration%time.Second != 0:
		return fmt.Errorf("the iperf time is whole seconds from 1s to %s (%s)", models.IperfMaxDuration, o.Duration)
	case o.Length < models.IperfMinLength || o.Length > models.IperfMaxLength:
		return fmt.Errorf("the iperf length is %d to %d octets (%d)", models.IperfMinLength, models.IperfMaxLength, o.Length)
	}
	return nil
}

// makeCookie returns a random NUL-terminated test cookie, as iperf3 makes it.
func makeCookie() []byte {
	b := make([]byte, cookieLen)
	for i := 0; i < cookieLen-1; i++ {
		b[i] = base32[rand.IntN(len(base32))]
	}
	return b // ends with NUL
}

// dataPacket builds a data datagram. The UDP checksum is left out (zero is
// allowed over IPv4) to save the work at high rates.
func dataPacket(src, dst netip.AddrPort, ipID uint16, payload []byte) []byte {
	b := make([]byte, 0, 28+len(payload))
	b = traffic.AppendIPv4(b, src.Addr(), dst.Addr(), models.ProtoUDP, ipID, false, 8+len(payload))
	b = binary.BigEndian.AppendUint16(b, src.Port())
	b = binary.BigEndian.AppendUint16(b, dst.Port())
	b = binary.BigEndian.AppendUint16(b, uint16(8+len(payload)))
	b = append(b, 0, 0)
	return append(b, payload...)
}

// newSender returns a sender of length-octet datagrams at rate bit/s.
func newSender(rate uint64, length int, send func([]byte) error) *sender {
	return &sender{rate: rate, length: length, send: send, payload: make([]byte, length)}
}

// pump sends what is due at now (at most a burst, after a stall).
func (s *sender) pump(now time.Time) {
	due := uint64(float64(s.rate) * now.Sub(s.start).Seconds() / 8)
	for n := 0; s.bytes < due && n < maxBurst; n++ {
		s.pcount++
		putHeader(s.payload, time.Now(), s.pcount)
		if err := s.send(s.payload); err != nil {
			s.pcount--
			s.errors++
			if s.firstErr == nil {
				s.firstErr = err
			}
			return
		}
		s.bytes += uint64(s.length)
		s.packets++
	}
}

// Run makes one UDP test from src against the server and returns the result.
func Run(ctx context.Context, st *traffic.Stack, link traffic.Link, src netip.Addr, o Options, report Reporter) (Result, error) {
	if err := o.Check(); err != nil {
		return Result{}, err
	}
	// The control connection, over the UE's small TCP.
	dctx, cancel := context.WithTimeout(ctx, connectTimeout)
	conn, err := tcp.Dial(dctx, st, link, src, o.Server)
	cancel()
	if err != nil {
		return Result{}, fmt.Errorf("cannot connect to %s (is iperf3 -s running?): %v", o.Server, err)
	}
	k := smallControl(conn)
	ok := false
	defer func() {
		if ok {
			conn.Close()
		} else {
			conn.Abort()
		}
	}()

	// The setup: cookie, parameters, then the server asks for the streams.
	if _, err := conn.Write(makeCookie()); err != nil {
		return Result{}, err
	}
	if err := k.expect(paramExchange, stepTimeout); err != nil {
		return Result{}, err
	}
	p := params{UDP: true, Time: int(o.Duration / time.Second), Parallel: 1, Reverse: o.Reverse,
		Len: o.Length, Bandwidth: o.Rate, PacingTimer: pacingTimer, ClientVersion: clientVersion}
	if err := k.writeJSON(p); err != nil {
		return Result{}, err
	}
	if err := k.expect(createStreams, stepTimeout); err != nil {
		return Result{}, err
	}

	// The data stream: our first datagram tells the server our address.
	recv := &receiver{}
	var connected atomic.Bool
	ready := make(chan struct{})
	port, err := st.Bind(models.ProtoUDP, 0, func(ip models.IPv4) {
		u, ok := traffic.ParseUDP(ip)
		if !ok || u.Src != o.Server {
			return
		}
		if !connected.Load() {
			if len(u.Payload) == udpConnectLen && connected.CompareAndSwap(false, true) {
				close(ready)
			}
			return
		}
		if o.Reverse {
			recv.take(u.Payload, time.Now())
		}
	})
	if err != nil {
		return Result{}, err
	}
	defer st.Unbind(models.ProtoUDP, port)
	local := netip.AddrPortFrom(src, port)
	if err := link.Send(traffic.BuildUDP(local, o.Server, st.IPID(), udpConnectMsg)); err != nil {
		return Result{}, fmt.Errorf("cannot send UDP: %v", err)
	}
	select {
	case <-ready:
	case <-time.After(stepTimeout):
		return Result{}, errors.New("no answer to the UDP connect (check the firewall and the UPF)")
	case <-ctx.Done():
		return Result{}, ctx.Err()
	}
	if err := k.expect(testStart, stepTimeout); err != nil {
		return Result{}, err
	}
	if err := k.expect(testRunning, stepTimeout); err != nil {
		return Result{}, err
	}

	// The test: send (uplink) or receive (downlink) until the time is up.
	snd := newSender(o.Rate, o.Length, func(payload []byte) error {
		return link.Send(dataPacket(local, o.Server, st.IPID(), payload))
	})
	start := time.Now()
	snd.start = start
	deadline := start.Add(o.Duration)
	tick := time.NewTicker(tickInterval)
	defer tick.Stop()
	second, lastBytes, last := 1, uint64(0), receiverSnapshot{}
	aborted := false
	for running := true; running; {
		select {
		case <-ctx.Done():
			aborted, running = true, false
			continue
		case now := <-tick.C:
			if !o.Reverse {
				if snd.pump(now); snd.errors == 1 && snd.firstErr != nil {
					report(true, fmt.Sprintf("cannot send: %v", snd.firstErr))
					snd.firstErr = nil
				}
			}
			// Any state from the server during the test ends it.
			if conn.Pending() {
				s, err := k.readState(promptReadTimeout)
				switch {
				case err != nil:
					return Result{}, fmt.Errorf("the control connection ended during the test: %v", err)
				case s == serverError:
					return Result{}, errors.New(k.serverErrorText())
				default:
					return Result{}, fmt.Errorf("the server sent %s during the test", stateName(s))
				}
			}
			// A progress line each second.
			if now.Sub(start) >= time.Duration(second)*time.Second {
				if o.Reverse {
					cur := recv.snapshot()
					n := cur.count - last.count
					lost := cur.lost - min(cur.lost, last.lost)
					loss := 0.0
					if n > 0 {
						loss = 100 * float64(lost) / float64(n)
					}
					report(false, fmt.Sprintf("%d–%d s: received %s, lost %d/%d (%.3g%%), jitter %.3f ms", second-1, second,
						traffic.FormatRate(float64(cur.bytes-last.bytes)*8), lost, n, loss, cur.jitter*1000))
					last = cur
				} else {
					report(false, fmt.Sprintf("%d–%d s: sent %s", second-1, second, traffic.FormatRate(float64(snd.bytes-lastBytes)*8)))
					lastBytes = snd.bytes
				}
				second++
			}
			if !now.Before(deadline) {
				running = false
			}
		}
	}
	elapsed := time.Since(start).Seconds()
	if aborted {
		k.writeState(clientTerminate)
		return Result{}, ErrAborted
	}

	// The end: exchange the results.
	if err := k.writeState(testEnd); err != nil {
		return Result{}, err
	}
	if err := k.expect(exchangeResults, resultsTimeout); err != nil {
		return Result{}, err
	}
	rs := recv.snapshot()
	mine := streamResult{ID: 1, Retransmits: -1, EndTime: elapsed}
	if o.Reverse {
		mine.Bytes, mine.Jitter, mine.Errors, mine.Packets = float64(rs.bytes), rs.jitter, float64(rs.lost), float64(rs.count)
	} else {
		mine.Bytes, mine.Packets = float64(snd.bytes), float64(snd.packets)
	}
	if err := k.writeJSON(results{Streams: []streamResult{mine}}); err != nil {
		return Result{}, err
	}
	var theirs results
	if err := k.readJSON(&theirs, resultsTimeout); err != nil {
		return Result{}, fmt.Errorf("cannot read the results of the server: %v", err)
	}
	if len(theirs.Streams) == 0 {
		return Result{}, errors.New("the results of the server have no stream")
	}
	if err := k.expect(displayResults, resultsTimeout); err != nil {
		return Result{}, err
	}
	k.writeState(iperfDone)
	ok = true

	// Combine our side with the server's: the sender counts, the receiver measures.
	server := theirs.Streams[0]
	r := Result{Reverse: o.Reverse, Seconds: elapsed}
	if o.Reverse {
		r.SentBytes, r.SentPackets = uint64(server.Bytes), uint64(server.Packets)
		r.RecvBytes, r.Lost, r.Jitter, r.OutOfOrder = rs.bytes, rs.lost, rs.jitter*1000, rs.outOfOrder
		r.Total = r.SentPackets
		if r.Total == 0 {
			r.Total = rs.count
		}
	} else {
		r.SentBytes, r.SentPackets = snd.bytes, snd.packets
		r.RecvBytes, r.Lost, r.Jitter = uint64(server.Bytes), uint64(server.Errors), server.Jitter*1000
		r.Total = snd.packets
	}
	return r, nil
}
