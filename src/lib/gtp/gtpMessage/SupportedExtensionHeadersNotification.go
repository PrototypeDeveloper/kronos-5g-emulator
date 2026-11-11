package gtpMessage

import (
	"fmt"
	"libgtp/gtpType"
)

// Extension Header Type List (141) uses an eight-bit length, unlike generic
// GTP TLVs. TS 29.281 7.2.3 and 8.5.
type SupportedExtensionHeadersNotification struct {
	Header gtpType.Header
	Types  []uint8
}

func NewSupportedExtensionHeadersNotification(sequence uint16, types []uint8) *SupportedExtensionHeadersNotification {
	h := gtpType.NewHeader(gtpType.MsgSupportedExtensionHeadersNotification, 0)
	h.S, h.SequenceNumber = true, sequence
	return &SupportedExtensionHeadersNotification{Header: h, Types: append([]uint8(nil), types...)}
}

func (m *SupportedExtensionHeadersNotification) Encode() ([]byte, error) {
	if m == nil || len(m.Types) == 0 || len(m.Types) > 255 {
		return nil, fmt.Errorf("invalid Extension Header Type List")
	}
	h := gtpType.NewHeader(gtpType.MsgSupportedExtensionHeadersNotification, 0)
	h.S, h.SequenceNumber, h.Length = true, m.Header.SequenceNumber, uint16(6+len(m.Types))
	b, err := h.EncodeMandatory()
	if err != nil {
		return nil, err
	}
	b = append(b, h.EncodeOptional()...)
	b = append(b, 141, byte(len(m.Types)))
	b = append(b, m.Types...)
	m.Header = h
	return b, nil
}

func DecodeSupportedExtensionHeadersNotification(data []byte) (*SupportedExtensionHeadersNotification, error) {
	h, err := gtpType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.MessageType != gtpType.MsgSupportedExtensionHeadersNotification || h.TEID != 0 || !h.S || h.E || h.PN {
		return nil, fmt.Errorf("invalid Supported Extension Headers Notification header")
	}
	if len(data) < 15 || data[12] != 141 || int(data[13]) != len(data)-14 {
		return nil, fmt.Errorf("invalid Extension Header Type List")
	}
	return &SupportedExtensionHeadersNotification{Header: h, Types: append([]uint8(nil), data[14:]...)}, nil
}
