package gtpType

import (
	"encoding/binary"
	"fmt"
)

const IEPrivateExtension uint8 = 255

// PrivateExtension is a TLV IE and may occur more than once in a signalling
// message.
type PrivateExtension struct {
	Header              IEHeader
	ExtensionIdentifier uint16
	ExtensionValue      []byte
}

func NewPrivateExtension(identifier uint16, value []byte) (PrivateExtension, error) {
	var p PrivateExtension
	if len(value)+2 > 0xffff {
		return p, fmt.Errorf("Private Extension too long: %d", len(value)+2)
	}
	p.Header = IEHeader{Type: IEPrivateExtension, Length: uint16(len(value) + 2)}
	p.ExtensionIdentifier = identifier
	p.ExtensionValue = append([]byte(nil), value...)
	return p, nil
}

func (p PrivateExtension) Encode() ([]byte, error) {
	if p.Header.Type != IEPrivateExtension {
		return nil, fmt.Errorf("Private Extension IE is not present")
	}
	valueLen := 2 + len(p.ExtensionValue)
	if valueLen > 0xffff {
		return nil, fmt.Errorf("Private Extension too long: %d", valueLen)
	}
	b := make([]byte, 3+valueLen)
	b[0] = IEPrivateExtension
	binary.BigEndian.PutUint16(b[1:3], uint16(valueLen))
	binary.BigEndian.PutUint16(b[3:5], p.ExtensionIdentifier)
	copy(b[5:], p.ExtensionValue)
	return b, nil
}

func DecodePrivateExtension(data []byte) (PrivateExtension, int, error) {
	var p PrivateExtension
	if len(data) < 5 {
		return p, 0, fmt.Errorf("Private Extension IE is truncated")
	}
	if data[0] != IEPrivateExtension {
		return p, 0, fmt.Errorf("unexpected Private Extension IE type: %d", data[0])
	}
	length := int(binary.BigEndian.Uint16(data[1:3]))
	if length < 2 {
		return p, 0, fmt.Errorf("Private Extension IE too short: %d", length)
	}
	total := 3 + length
	if len(data) < total {
		return p, 0, fmt.Errorf("Private Extension IE length exceeds packet: %d", length)
	}
	p.Header = IEHeader{Type: IEPrivateExtension, Length: uint16(length)}
	p.ExtensionIdentifier = binary.BigEndian.Uint16(data[3:5])
	p.ExtensionValue = append([]byte(nil), data[5:total]...)
	return p, total, nil
}
