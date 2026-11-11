// Package traffic is the user-plane traffic of the emulated UEs: a small
// IPv4 stack per UE that sends and receives the packets of
// ping and iperf through the PDU session's GTP-U tunnel. The engine passes
// the downlink packets in (Stack.Deliver) and gives a Link that sends uplink.
// Nothing here needs privileges: the packets never touch the host's stack.
package traffic

import (
	"encoding/binary"
	"kronos/models"
	"net/netip"
)

// Link sends one IPv4 packet uplink through the UE's PDU session.
type Link interface {
	Send(pkt []byte) error
}

// LinkFunc adapts a function to Link.
type LinkFunc func(pkt []byte) error

// Send implements Link.
func (f LinkFunc) Send(pkt []byte) error { return f(pkt) }

// ParseIPv4 reads the header of an IPv4 packet. Padding after the total
// length is dropped; a bad header (length, checksum) is refused.
func ParseIPv4(pkt []byte) (models.IPv4, bool) {
	if len(pkt) < 20 || pkt[0]>>4 != 4 {
		return models.IPv4{}, false
	}
	ihl := int(pkt[0]&0x0f) * 4
	total := int(binary.BigEndian.Uint16(pkt[2:]))
	if ihl < 20 || total < ihl || total > len(pkt) || Checksum(pkt[:ihl], 0) != 0 {
		return models.IPv4{}, false
	}
	return models.IPv4{
		Src:     netip.AddrFrom4([4]byte(pkt[12:16])),
		Dst:     netip.AddrFrom4([4]byte(pkt[16:20])),
		Proto:   pkt[9],
		TTL:     pkt[8],
		ID:      binary.BigEndian.Uint16(pkt[4:]),
		Frag:    binary.BigEndian.Uint16(pkt[6:])&0x3fff != 0,
		Payload: pkt[ihl:total],
	}, true
}

// AppendIPv4 appends an IPv4 header (20 octets, checksum set) for a payload
// of n octets. df sets "don't fragment".
func AppendIPv4(b []byte, src, dst netip.Addr, proto uint8, id uint16, df bool, n int) []byte {
	h := make([]byte, 20)
	h[0] = 0x45
	binary.BigEndian.PutUint16(h[2:], uint16(20+n))
	binary.BigEndian.PutUint16(h[4:], id)
	if df {
		h[6] = 0x40
	}
	h[8] = defaultTTL // TTL
	h[9] = proto
	s4, d4 := src.As4(), dst.As4()
	copy(h[12:], s4[:])
	copy(h[16:], d4[:])
	binary.BigEndian.PutUint16(h[10:], Checksum(h, 0))
	return append(b, h...)
}

// Checksum is the Internet checksum (RFC 1071) of b added to a partial sum.
func Checksum(b []byte, sum uint32) uint16 {
	for len(b) >= 2 {
		sum += uint32(b[0])<<8 | uint32(b[1])
		b = b[2:]
	}
	if len(b) == 1 {
		sum += uint32(b[0]) << 8
	}
	for sum>>16 != 0 {
		sum = sum&0xffff + sum>>16
	}
	return ^uint16(sum)
}

// pseudoSum is the partial sum of the IPv4 pseudo-header of TCP and UDP.
func pseudoSum(src, dst netip.Addr, proto uint8, length int) uint32 {
	s4, d4 := src.As4(), dst.As4()
	var sum uint32
	for i := 0; i < 4; i += 2 {
		sum += uint32(s4[i])<<8 | uint32(s4[i+1])
		sum += uint32(d4[i])<<8 | uint32(d4[i+1])
	}
	return sum + uint32(proto) + uint32(length)
}

// TransportChecksum is the TCP or UDP checksum of seg (whose own checksum
// field is zero when computing, or set when verifying: 0 means valid).
func TransportChecksum(src, dst netip.Addr, proto uint8, seg []byte) uint16 {
	return Checksum(seg, pseudoSum(src, dst, proto, len(seg)))
}

// BuildUDP returns an IPv4 packet that carries one UDP datagram.
func BuildUDP(src, dst netip.AddrPort, id uint16, payload []byte) []byte {
	b := make([]byte, 0, 28+len(payload))
	b = AppendIPv4(b, src.Addr(), dst.Addr(), models.ProtoUDP, id, false, 8+len(payload))
	b = append(b, 0, 0, 0, 0, 0, 0, 0, 0)
	u := b[20:28]
	binary.BigEndian.PutUint16(u[0:], src.Port())
	binary.BigEndian.PutUint16(u[2:], dst.Port())
	binary.BigEndian.PutUint16(u[4:], uint16(8+len(payload)))
	b = append(b, payload...)
	c := TransportChecksum(src.Addr(), dst.Addr(), models.ProtoUDP, b[20:])
	if c == 0 {
		c = 0xffff // zero means "no checksum" in UDP
	}
	binary.BigEndian.PutUint16(b[26:], c)
	return b
}

// ParseUDP reads the UDP datagram of an IPv4 packet (checksum verified when set).
func ParseUDP(ip models.IPv4) (models.UDP, bool) {
	p := ip.Payload
	if ip.Proto != models.ProtoUDP || ip.Frag || len(p) < 8 {
		return models.UDP{}, false
	}
	n := int(binary.BigEndian.Uint16(p[4:]))
	if n < 8 || n > len(p) {
		return models.UDP{}, false
	}
	if binary.BigEndian.Uint16(p[6:]) != 0 && TransportChecksum(ip.Src, ip.Dst, models.ProtoUDP, p[:n]) != 0 {
		return models.UDP{}, false
	}
	return models.UDP{
		Src:     netip.AddrPortFrom(ip.Src, binary.BigEndian.Uint16(p[0:])),
		Dst:     netip.AddrPortFrom(ip.Dst, binary.BigEndian.Uint16(p[2:])),
		Payload: p[8:n],
	}, true
}

// BuildEcho returns an IPv4 packet with an ICMP echo request (typ 8) or
// reply (typ 0) carrying data.
func BuildEcho(src, dst netip.Addr, ipID uint16, typ uint8, id, seq uint16, data []byte) []byte {
	b := make([]byte, 0, 28+len(data))
	df := 28+len(data) <= maxDFEcho // larger ones may be fragmented on the way
	b = AppendIPv4(b, src, dst, models.ProtoICMP, ipID, df, 8+len(data))
	b = append(b, typ, 0, 0, 0, byte(id>>8), byte(id), byte(seq>>8), byte(seq))
	b = append(b, data...)
	binary.BigEndian.PutUint16(b[22:], Checksum(b[20:], 0))
	return b
}

// ParseEcho reads an ICMP echo request or reply (checksum verified).
func ParseEcho(ip models.IPv4) (models.Echo, bool) {
	p := ip.Payload
	if ip.Proto != models.ProtoICMP || ip.Frag || len(p) < 8 || (p[0] != icmpEchoReply && p[0] != icmpEchoRequest) || p[1] != 0 || Checksum(p, 0) != 0 {
		return models.Echo{}, false
	}
	return models.Echo{Type: p[0], ID: binary.BigEndian.Uint16(p[4:]), Seq: binary.BigEndian.Uint16(p[6:]), Data: p[8:]}, true
}
