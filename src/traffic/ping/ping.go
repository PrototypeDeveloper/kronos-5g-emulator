// Package ping sends ICMP echo requests from a UE through its PDU session
// and measures the replies, like the ping of Linux: requests
// go out at a fixed interval whatever the replies do, and the summary gives
// the loss and min/avg/max/mdev of the round-trip times.
package ping

import (
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"net/netip"
	"time"

	"kronos/models"
	"kronos/traffic"
)

// Avg and Mdev are computed like the ping of Linux (mdev: the standard
// deviation of the RTTs).
func (s Stats) Avg() float64 {
	if s.Received == 0 {
		return 0
	}
	return s.sum / float64(s.Received)
}

// Mdev is the standard deviation of the RTTs in milliseconds.
func (s Stats) Mdev() float64 {
	if s.Received == 0 {
		return 0
	}
	avg := s.Avg()
	return math.Sqrt(math.Max(0, s.sumSquares/float64(s.Received)-avg*avg))
}

// ms converts d to milliseconds.
func ms(d time.Duration) float64 { return float64(d) / float64(time.Millisecond) }

// String renders the summary: 5 transmitted, 5 received, 0% loss, time 4005 ms,
// RTT min/avg/max/mdev = 0.218/0.251/0.298/0.031 ms
func (s Stats) String() string {
	loss := 100
	if s.Sent > 0 {
		loss = (s.Sent - s.Received) * 100 / s.Sent
	}
	out := fmt.Sprintf("%d transmitted, %d received, %d%% loss, time %d ms", s.Sent, s.Received, loss, s.Elapsed.Milliseconds())
	if s.Received > 0 {
		out += fmt.Sprintf(", rtt min/avg/max/mdev = %.3f/%.3f/%.3f/%.3f ms", ms(s.Min), s.Avg(), ms(s.Max), s.Mdev())
	}
	if s.SendErrors > 0 {
		out += fmt.Sprintf(", %d could not be sent", s.SendErrors)
	}
	return out
}

// add counts a reply with its round-trip time.
func (s *Stats) add(rtt time.Duration) {
	if s.Received == 0 || rtt < s.Min {
		s.Min = rtt
	}
	if rtt > s.Max {
		s.Max = rtt
	}
	s.Received++
	m := ms(rtt)
	s.sum += m
	s.sumSquares += m * m
}

// Reporter receives the progress lines; warn marks a lost reply or an error.
type Reporter func(warn bool, text string)

// Run pings from src until Count requests are answered or have timed out, or
// ctx ends (then the requests still in flight count as lost).
func Run(ctx context.Context, st *traffic.Stack, link traffic.Link, src netip.Addr, o models.PingOptions, report Reporter) (Stats, error) {
	if o.Interval <= 0 {
		o.Interval = DefaultInterval
	}
	if o.Size < 0 {
		o.Size = DefaultSize
	}
	switch {
	case o.Interval < MinInterval:
		return Stats{}, fmt.Errorf("the ping interval must be at least %s", MinInterval)
	case o.Size > MaxSize:
		return Stats{}, fmt.Errorf("the ping size is 0 to %d", MaxSize)
	case o.Count < 0:
		return Stats{}, fmt.Errorf("the ping count must be 0 or more (0: until stop)")
	}
	// The replies come on an ICMP identifier of our own.
	replies := make(chan reply, replyQueueLen)
	id, err := st.Bind(models.ProtoICMP, 0, func(ip models.IPv4) {
		if e, ok := traffic.ParseEcho(ip); ok && e.Type == icmpEchoReply && ip.Src == o.Dst {
			select {
			case replies <- reply{e.Seq, time.Now()}:
			default:
			}
		}
	})
	if err != nil {
		return Stats{}, err
	}
	defer st.Unbind(models.ProtoICMP, id)

	verbose := o.Interval >= QuietBelow
	data := make([]byte, o.Size)
	for i := range data {
		data[i] = byte(i)
	}
	var stats Stats
	pending := map[uint16]time.Time{}
	var seq uint16
	start := time.Now()
	failing := false
	// send sends the next request; the payload starts with the send time.
	send := func() {
		seq++
		if len(data) >= 8 {
			binary.BigEndian.PutUint64(data, uint64(time.Now().UnixNano()))
		}
		at := time.Now()
		if err := link.Send(traffic.BuildEcho(src, o.Dst, st.IPID(), icmpEchoRequest, id, seq, data)); err != nil {
			stats.SendErrors++
			if !failing {
				report(true, fmt.Sprintf("seq=%d cannot be sent: %v", seq, err))
			}
			failing = true
			return
		}
		if failing {
			report(false, fmt.Sprintf("sending again from seq=%d", seq))
			failing = false
		}
		stats.Sent++
		pending[seq] = at
	}
	sending := func() bool { return o.Count == 0 || stats.Sent+stats.SendErrors < o.Count }

	// The loop: a request each interval, replies matched by seq, timeouts checked.
	send()
	tick := time.NewTicker(o.Interval)
	defer tick.Stop()
	check := time.NewTicker(min(checkInterval, o.Interval))
	defer check.Stop()
	progress := time.NewTicker(time.Second)
	defer progress.Stop()
	stopped := ctx.Done()
	var grace <-chan time.Time
	for sending() || len(pending) > 0 {
		select {
		case <-stopped:
			// Stopped: no more requests; the ones in flight get a moment
			// to be answered, so stopping does not count them as lost.
			o.Count = stats.Sent + stats.SendErrors
			stopped = nil
			grace = time.After(stopGrace)
		case <-grace:
			stats.Elapsed = time.Since(start)
			return stats, nil
		case <-tick.C:
			if sending() {
				send()
			}
		case r := <-replies:
			sent, ok := pending[r.seq]
			if !ok {
				continue // late (already counted lost) or duplicate
			}
			delete(pending, r.seq)
			rtt := r.at.Sub(sent)
			stats.add(rtt)
			if verbose {
				report(false, fmt.Sprintf("seq=%d reply %.3f ms", r.seq, ms(rtt)))
			}
		case now := <-check.C:
			for s, sent := range pending {
				if now.Sub(sent) > timeout {
					delete(pending, s)
					if verbose {
						report(true, fmt.Sprintf("seq=%d no reply (%s)", s, timeout))
					}
				}
			}
		case <-progress.C:
			if !verbose {
				report(false, fmt.Sprintf("%d s: %d transmitted, %d received", int(time.Since(start).Seconds()+0.5), stats.Sent, stats.Received))
			}
		}
	}
	stats.Elapsed = time.Since(start)
	return stats, nil
}
