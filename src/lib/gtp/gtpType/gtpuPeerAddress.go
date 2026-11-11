package gtpType

import (
	"encoding/binary"
	"fmt"
	"net"
)

const IEGTPUPeerAddress uint8 = 133

// GTPUPeerAddress is a TLV IE containing exactly one IPv4 or IPv6 address.
// Header.Length is 4 for IPv4 and 16 for IPv6.
type GTPUPeerAddress struct {
	Header IEHeader
	IPv4   [4]byte
	IPv6   [16]byte
}

func NewGTPUPeerAddress(ip net.IP) (GTPUPeerAddress, error) {
	var v GTPUPeerAddress
	if ip == nil {
		return v, fmt.Errorf("GTP-U Peer Address is nil")
	}
	if ip4 := ip.To4(); ip4 != nil {
		v.Header = IEHeader{Type: IEGTPUPeerAddress, Length: 4}
		copy(v.IPv4[:], ip4)
		return v, nil
	}
	ip16 := ip.To16()
	if ip16 == nil {
		return v, fmt.Errorf("invalid GTP-U Peer Address: %v", ip)
	}
	v.Header = IEHeader{Type: IEGTPUPeerAddress, Length: 16}
	copy(v.IPv6[:], ip16)
	return v, nil
}

func (v GTPUPeerAddress) IsIPv4() bool {
	return v.Header.Type == IEGTPUPeerAddress && v.Header.Length == 4
}
func (v GTPUPeerAddress) IsIPv6() bool {
	return v.Header.Type == IEGTPUPeerAddress && v.Header.Length == 16
}

func (v GTPUPeerAddress) IP() net.IP {
	if v.IsIPv4() {
		return net.IPv4(v.IPv4[0], v.IPv4[1], v.IPv4[2], v.IPv4[3])
	}
	if v.IsIPv6() {
		return net.IP(append([]byte(nil), v.IPv6[:]...))
	}
	return nil
}

func (v GTPUPeerAddress) Encode() ([]byte, error) {
	if v.Header.Type != IEGTPUPeerAddress {
		return nil, fmt.Errorf("GTP-U Peer Address IE is not present")
	}
	var value []byte
	switch v.Header.Length {
	case 4:
		value = v.IPv4[:]
	case 16:
		value = v.IPv6[:]
	default:
		return nil, fmt.Errorf("GTP-U Peer Address length must be 4 or 16: %d", v.Header.Length)
	}
	b := make([]byte, 3+len(value))
	b[0] = IEGTPUPeerAddress
	binary.BigEndian.PutUint16(b[1:3], uint16(len(value)))
	copy(b[3:], value)
	return b, nil
}

func DecodeGTPUPeerAddress(data []byte) (GTPUPeerAddress, int, error) {
	var v GTPUPeerAddress
	if len(data) < 3 {
		return v, 0, fmt.Errorf("GTP-U Peer Address IE is truncated")
	}
	if data[0] != IEGTPUPeerAddress {
		return v, 0, fmt.Errorf("unexpected GTP-U Peer Address IE type: %d", data[0])
	}
	length := int(binary.BigEndian.Uint16(data[1:3]))
	if length != 4 && length != 16 {
		return v, 0, fmt.Errorf("GTP-U Peer Address length must be 4 or 16: %d", length)
	}
	total := 3 + length
	if len(data) < total {
		return v, 0, fmt.Errorf("GTP-U Peer Address IE length exceeds packet: %d", length)
	}
	v.Header = IEHeader{Type: IEGTPUPeerAddress, Length: uint16(length)}
	if length == 4 {
		copy(v.IPv4[:], data[3:7])
	} else {
		copy(v.IPv6[:], data[3:19])
	}
	return v, total, nil
}
