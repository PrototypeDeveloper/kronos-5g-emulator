package gtpType

import (
	"encoding/binary"
	"fmt"
)

// IEHeader represents GTPv1-U Information Element metadata.
// Length is the value length. For TV IEs it is not carried on the wire but is
// populated with the fixed value length so decoded structures remain easy to inspect.
type IEHeader struct {
	Type   uint8
	Length uint16
}

// InformationElement preserves an unimplemented TLV Information Element.
// Unknown TV IEs cannot be preserved generically because their value length is
// defined by the individual IE type.
type InformationElement struct {
	Header IEHeader
	Value  []byte
}

func NewInformationElement(ieType uint8, value []byte) (InformationElement, error) {
	var ie InformationElement
	if ieType < 128 {
		return ie, fmt.Errorf("generic InformationElement supports TLV IE types only: %d", ieType)
	}
	if len(value) > 0xffff {
		return ie, fmt.Errorf("information element too long: %d", len(value))
	}
	ie.Header = IEHeader{Type: ieType, Length: uint16(len(value))}
	ie.Value = append([]byte(nil), value...)
	return ie, nil
}

func (ie InformationElement) Encode() ([]byte, error) {
	if ie.Header.Type < 128 {
		return nil, fmt.Errorf("generic InformationElement supports TLV IE types only: %d", ie.Header.Type)
	}
	if len(ie.Value) > 0xffff {
		return nil, fmt.Errorf("information element too long: %d", len(ie.Value))
	}
	b := make([]byte, 3+len(ie.Value))
	b[0] = ie.Header.Type
	binary.BigEndian.PutUint16(b[1:3], uint16(len(ie.Value)))
	copy(b[3:], ie.Value)
	return b, nil
}

// DecodeTLVInformationElement decodes one TLV IE beginning at data[0].
func DecodeTLVInformationElement(data []byte) (InformationElement, int, error) {
	var ie InformationElement
	if len(data) < 3 {
		return ie, 0, fmt.Errorf("TLV information element is truncated")
	}
	if data[0] < 128 {
		return ie, 0, fmt.Errorf("IE type %d is not TLV", data[0])
	}
	length := int(binary.BigEndian.Uint16(data[1:3]))
	total := 3 + length
	if len(data) < total {
		return ie, 0, fmt.Errorf("TLV information element length exceeds packet: type=%d length=%d", data[0], length)
	}
	ie.Header = IEHeader{Type: data[0], Length: uint16(length)}
	ie.Value = append([]byte(nil), data[3:total]...)
	return ie, total, nil
}
