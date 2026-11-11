package engine

import (
	"context"
	"errors"
	"fmt"
	"net/netip"

	"kronos/models"
	"kronos/traffic"
	"kronos/traffic/iperf"
	"kronos/traffic/ping"
)

// The user-plane traffic of a subscriber: ping and iperf run
// on goroutines of their own, through the subscriber's IP stack and its PDU
// session; their lines and results arrive as events. stop, a reset and the
// end of the engine stop them.

// session picks the PDU session with the lowest ID whose tunnel is up (on the actor).
func (s *subscriber) session() (models.Tunnel, error) {
	var t models.Tunnel
	found := false
	for id, x := range s.tunnels {
		if !found || id < t.SessionID {
			t, found = x, true
		}
	}
	if !found {
		return t, errors.New("no usable PDU session (establish one first; an idle UE comes back with a Service Request)")
	}
	return t, nil
}

// link sends uplink through the session's current tunnel.
func (s *subscriber) link(session uint8) traffic.Link {
	return traffic.LinkFunc(func(pkt []byte) error { return s.e.uplane.sendFor(s, session, pkt) })
}

// startRun runs body on its own goroutine until it ends or is stopped. The
// lines it reports and its summary become events of the subscriber.
func (s *subscriber) startRun(kind, label string, body func(ctx context.Context, report func(warn bool, text string)) (string, error)) {
	ctx, cancel := context.WithCancel(s.e.runCtx())
	s.runs.mu.Lock()
	id := s.runs.next
	s.runs.next++
	s.runs.runs[id] = trafficRun{kind: kind, cancel: cancel}
	s.runs.mu.Unlock()
	publish := func(kind, text string) {
		s.e.publish(models.Event{Kind: kind, Source: s.imsi, Text: label + ": " + text})
	}
	report := func(warn bool, text string) {
		if warn {
			publish(models.KindWarn, text)
		} else {
			publish(models.KindInfo, text)
		}
	}
	go func() {
		defer func() {
			s.runs.mu.Lock()
			delete(s.runs.runs, id)
			s.runs.mu.Unlock()
			cancel()
		}()
		summary, err := body(ctx, report)
		switch {
		case errors.Is(err, iperf.ErrAborted):
			publish(models.KindInfo, "stopped")
		case err != nil:
			publish(models.KindError, err.Error())
		case summary != "":
			publish(models.KindInfo, summary)
		}
	}()
}

// stopRuns stops the runs of a kind ("" for all); they report what they have.
func (s *subscriber) stopRuns(kind string) int {
	s.runs.mu.Lock()
	defer s.runs.mu.Unlock()
	n := 0
	for _, r := range s.runs.runs {
		if kind == "" || r.kind == kind {
			r.cancel()
			n++
		}
	}
	return n
}

// Ping implements command.Env: ICMP echo through the PDU session with the
// lowest ID (traffic/ping).
func (s *subscriber) Ping(a models.PingArgs) error {
	dst, err := netip.ParseAddr(a.Dst)
	if err != nil || !dst.Is4() {
		return fmt.Errorf("the ping destination %q is not an IPv4 address", a.Dst)
	}
	t, err := s.session()
	if err != nil {
		return err
	}
	link := s.link(t.SessionID)
	o := models.PingOptions{Dst: dst, Count: a.Count, Interval: a.Interval, Size: a.Size}
	s.startRun("ping", "ping "+a.Dst, func(ctx context.Context, report func(bool, string)) (string, error) {
		st, err := ping.Run(ctx, s.stack, link, t.UEIP, o, report)
		if err != nil {
			return "", err
		}
		return st.String(), nil
	})
	return nil
}

// Iperf implements command.Env: an iperf3 UDP test through the PDU session
// with the lowest ID (traffic/iperf), or a TCP test (iperftcp) through the
// subscriber's TUN device.
func (s *subscriber) Iperf(a models.IperfArgs) error {
	dst, err := netip.ParseAddr(a.Dst)
	if err != nil || !dst.Is4() {
		return fmt.Errorf("the iperf destination %q is not an IPv4 address", a.Dst)
	}
	t, err := s.session()
	if err != nil {
		return err
	}
	if a.TCP {
		return s.iperfTCP(a, dst, t)
	}
	o := iperf.Options{Server: netip.AddrPortFrom(dst, uint16(a.Port)), Rate: a.Rate, Duration: a.Duration, Reverse: a.Reverse, Length: a.Length}
	if err := o.Check(); err != nil {
		return err
	}
	link := s.link(t.SessionID)
	label := "iperf " + o.Server.String()
	dir := "uplink (UE → DN)"
	if o.Reverse {
		dir = "downlink (DN → UE)"
	}
	s.e.publish(models.Event{Kind: models.KindInfo, Source: s.imsi, Text: fmt.Sprintf("%s: %s %s × %s, %d octets, starting (UE %s)",
		label, dir, traffic.FormatRate(float64(o.Rate)), o.Duration, o.Length, t.UEIP)})
	s.startRun("iperf", label, func(ctx context.Context, report func(bool, string)) (string, error) {
		r, err := iperf.Run(ctx, s.stack, link, t.UEIP, o, report)
		if err != nil {
			return "", err
		}
		return r.String(), nil
	})
	return nil
}

// iperfTCP runs a TCP test with host sockets bound to the UE address,
// which the subscriber's TUN device carries through its PDU session.
func (s *subscriber) iperfTCP(a models.IperfArgs, dst netip.Addr, t models.Tunnel) error {
	x := s.tun.Load()
	switch {
	case x == nil && s.e.openTUN == nil:
		return errors.New("iperftcp runs through the UE's TUN device: set tun: true under ue: in the config (gnodeb needs CAP_NET_ADMIN)")
	case x == nil:
		return errors.New("the UE has no TUN device (see the TUN error when the session came up)")
	case x.ueIP != t.UEIP:
		return fmt.Errorf("the TUN device carries %s (PDU session %d), not the session in use (%s)", x.ueIP, x.session, t.UEIP)
	}
	o := iperf.TCPOptions{Server: netip.AddrPortFrom(dst, uint16(a.Port)), Local: x.ueIP, Rate: a.Rate, Duration: a.Duration, Reverse: a.Reverse, Length: a.Length}
	if err := o.Check(); err != nil {
		return err
	}
	label := "iperftcp " + o.Server.String()
	dir := "uplink (UE → DN)"
	if o.Reverse {
		dir = "downlink (DN → UE)"
	}
	rate := "as fast as TCP goes"
	if o.Rate > 0 {
		rate = traffic.FormatRate(float64(o.Rate))
	}
	s.e.publish(models.Event{Kind: models.KindInfo, Source: s.imsi, Text: fmt.Sprintf("%s: TCP %s, %s × %s, %d-octet writes, starting (UE %s, %s)",
		label, dir, rate, o.Duration, o.Length, x.ueIP, x.dev.Name())})
	s.startRun("iperf", label, func(ctx context.Context, report func(bool, string)) (string, error) {
		r, err := iperf.RunTCP(ctx, o, report)
		if err != nil {
			return "", err
		}
		return r.String(), nil
	})
	return nil
}

// Stop implements command.Env.
func (s *subscriber) Stop(kind string) error {
	s.stopRuns(kind)
	return nil
}
