package traffic

import (
	"errors"
	"fmt"
	"kronos/models"
	"math/rand/v2"
)

// Handler takes a downlink packet for an endpoint. It runs on the goroutine
// that delivers (the GTP-U reader), so it must be quick and must copy what
// it keeps: the packet's memory is reused.
type Handler func(ip models.IPv4)

// NewStack returns an empty endpoint table. Its ephemeral ports start at a
// random place (RFC 6056), so a new process does not take the ports of the
// last one again: a server may still hold those connections (TIME_WAIT).
func NewStack() *Stack {
	return &Stack{handlers: map[endpoint]Handler{}, next: map[uint8]uint16{
		models.ProtoTCP: ephemeralLo + uint16(rand.IntN(ephemeralCount)),
		models.ProtoUDP: ephemeralLo + uint16(rand.IntN(ephemeralCount)),
	}}
}

// ErrInUse is returned when an endpoint is taken.
var ErrInUse = errors.New("port in use")

// Bind gives the packets of proto to port to h. Port 0 takes a free
// ephemeral port (49152–65535; any identifier for ICMP), which is returned.
func (s *Stack) Bind(proto uint8, port uint16, h Handler) (uint16, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if port != 0 {
		if _, taken := s.handlers[endpoint{proto, port}]; taken {
			return 0, fmt.Errorf("%w (%d)", ErrInUse, port)
		}
		s.handlers[endpoint{proto, port}] = h
		return port, nil
	}
	// Port 0: the next free one after the last given out, wrapping around.
	lo := uint16(ephemeralLo)
	if proto == models.ProtoICMP {
		lo = 1
	}
	n := int(maxPort-lo) + 1
	start := s.next[proto] + 1 // wraps to 0 after 65535
	if start < lo {
		start = lo
	}
	for i := 0; i < n; i++ {
		p := lo + uint16((int(start-lo)+i)%n)
		if _, taken := s.handlers[endpoint{proto, p}]; !taken {
			s.handlers[endpoint{proto, p}] = h
			s.next[proto] = p
			return p, nil
		}
	}
	return 0, errors.New("no free port")
}

// Unbind frees an endpoint.
func (s *Stack) Unbind(proto uint8, port uint16) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.handlers, endpoint{proto, port})
}

// IPID returns the next IPv4 identification for an uplink packet.
func (s *Stack) IPID() uint16 { return uint16(s.ipID.Add(1)) }

// Unknown counts the downlink packets that no endpoint wanted.
func (s *Stack) Unknown() uint64 { return s.unknown.Load() }

// Deliver hands a downlink IPv4 packet to its endpoint: UDP and TCP by
// destination port, ICMP echo replies by identifier. It reports whether an
// endpoint took the packet.
func (s *Stack) Deliver(pkt []byte) bool {
	ip, ok := ParseIPv4(pkt)
	if !ok || ip.Frag {
		s.unknown.Add(1)
		return false
	}
	var ep endpoint
	switch p := ip.Payload; ip.Proto {
	case models.ProtoUDP, models.ProtoTCP:
		if len(p) < 4 {
			s.unknown.Add(1)
			return false
		}
		ep = endpoint{ip.Proto, uint16(p[2])<<8 | uint16(p[3])}
	case models.ProtoICMP:
		if len(p) < 8 || p[0] != icmpEchoReply { // only echo replies have an owner
			s.unknown.Add(1)
			return false
		}
		ep = endpoint{models.ProtoICMP, uint16(p[4])<<8 | uint16(p[5])}
	default:
		s.unknown.Add(1)
		return false
	}
	s.mu.Lock()
	h := s.handlers[ep]
	s.mu.Unlock()
	if h == nil {
		s.unknown.Add(1)
		return false
	}
	h(ip)
	return true
}
