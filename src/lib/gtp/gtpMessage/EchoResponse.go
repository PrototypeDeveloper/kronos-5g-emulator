package gtpMessage

import (
	"fmt"
	"sort"

	"libgtp/gtpType"
)

// EchoResponse is the response to an EchoRequest. Recovery is mandatory.
type EchoResponse struct {
	Header            gtpType.Header
	Recovery          gtpType.Recovery
	RecoveryTimeStamp gtpType.RecoveryTimeStamp
	PrivateExtension  []gtpType.PrivateExtension
	AdditionalIEs     []gtpType.InformationElement
}

func NewEchoResponse(sequenceNumber uint16, recovery gtpType.Recovery) *EchoResponse {
	h := gtpType.NewHeader(gtpType.MsgEchoResponse, 0)
	h.S = true
	h.SequenceNumber = sequenceNumber
	return &EchoResponse{
		Header:   h,
		Recovery: recovery,
	}
}

func (m *EchoResponse) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("Echo Response is nil")
	}
	if m.Recovery.Header.Type != gtpType.IERecovery {
		return nil, fmt.Errorf("mandatory Recovery IE is missing")
	}

	h := m.Header
	h.Version = gtpType.Version1
	h.ProtocolType = true
	h.MessageType = gtpType.MsgEchoResponse
	h.TEID = 0
	h.S = true
	h.E = false
	h.PN = false
	h.NPDUNumber = 0
	h.NextExtensionHeaderType = 0

	body := make([]byte, 0)
	b, err := m.Recovery.Encode()
	if err != nil {
		return nil, err
	}
	body = append(body, b...) // TV IE type 14 is always first by ascending type.

	type encodedIE struct {
		typeID uint8
		data   []byte
	}
	var ies []encodedIE
	if m.RecoveryTimeStamp.Header.Type != 0 {
		b, err := m.RecoveryTimeStamp.Encode()
		if err != nil {
			return nil, err
		}
		ies = append(ies, encodedIE{typeID: gtpType.IERecoveryTimeStamp, data: b})
	}
	for _, p := range m.PrivateExtension {
		b, err := p.Encode()
		if err != nil {
			return nil, err
		}
		ies = append(ies, encodedIE{typeID: gtpType.IEPrivateExtension, data: b})
	}
	for _, ie := range m.AdditionalIEs {
		b, err := ie.Encode()
		if err != nil {
			return nil, err
		}
		ies = append(ies, encodedIE{typeID: ie.Header.Type, data: b})
	}
	sort.SliceStable(ies, func(i, j int) bool { return ies[i].typeID < ies[j].typeID })
	for _, ie := range ies {
		body = append(body, ie.data...)
	}

	payloadLen := 4 + len(body)
	if payloadLen > 0xffff {
		return nil, fmt.Errorf("Echo Response too large: %d", payloadLen)
	}
	h.Length = uint16(payloadLen)

	mandatory, err := h.EncodeMandatory()
	if err != nil {
		return nil, err
	}
	out := make([]byte, 0, 8+payloadLen)
	out = append(out, mandatory...)
	out = append(out, h.EncodeOptional()...)
	out = append(out, body...)
	m.Header = h
	return out, nil
}

func DecodeEchoResponse(data []byte) (*EchoResponse, error) {
	h, err := gtpType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.MessageType != gtpType.MsgEchoResponse {
		return nil, fmt.Errorf("unexpected GTP message type: %d", h.MessageType)
	}
	if h.TEID != 0 {
		return nil, fmt.Errorf("Echo Response TEID must be zero: 0x%08x", h.TEID)
	}
	if !h.S {
		return nil, fmt.Errorf("Echo Response S flag must be set")
	}
	if h.E || h.PN {
		return nil, fmt.Errorf("unsupported Echo Response header flags: E=%v PN=%v", h.E, h.PN)
	}

	m := &EchoResponse{Header: h}
	offset := 12
	for offset < len(data) {
		ieType := data[offset]
		switch ieType {
		case gtpType.IERecovery:
			if m.Recovery.Header.Type != 0 {
				return nil, fmt.Errorf("duplicate Recovery IE")
			}
			v, n, err := gtpType.DecodeRecovery(data[offset:])
			if err != nil {
				return nil, err
			}
			m.Recovery = v
			offset += n
		case gtpType.IERecoveryTimeStamp:
			if m.RecoveryTimeStamp.Header.Type != 0 {
				return nil, fmt.Errorf("duplicate Recovery Time Stamp IE")
			}
			v, n, err := gtpType.DecodeRecoveryTimeStamp(data[offset:])
			if err != nil {
				return nil, err
			}
			m.RecoveryTimeStamp = v
			offset += n
		case gtpType.IEPrivateExtension:
			v, n, err := gtpType.DecodePrivateExtension(data[offset:])
			if err != nil {
				return nil, err
			}
			m.PrivateExtension = append(m.PrivateExtension, v)
			offset += n
		default:
			if ieType < 128 {
				return nil, fmt.Errorf("unsupported TV IE in Echo Response: %d", ieType)
			}
			v, n, err := gtpType.DecodeTLVInformationElement(data[offset:])
			if err != nil {
				return nil, err
			}
			m.AdditionalIEs = append(m.AdditionalIEs, v)
			offset += n
		}
	}
	if m.Recovery.Header.Type != gtpType.IERecovery {
		return nil, fmt.Errorf("mandatory Recovery IE is missing")
	}
	return m, nil
}
