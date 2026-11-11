package gtpMessage

import (
	"fmt"
	"sort"

	"libgtp/gtpType"
)

// ErrorIndication is sent when a G-PDU is received for an unknown tunnel.
// The GTP header TEID is always zero. The offending tunnel is identified by
// TunnelEndpointIdentifierDataI together with GTPUPeerAddress.
type ErrorIndication struct {
	Header gtpType.Header

	TunnelEndpointIdentifierDataI gtpType.TunnelEndpointIdentifierDataI
	GTPUPeerAddress               gtpType.GTPUPeerAddress
	RecoveryTimeStamp             gtpType.RecoveryTimeStamp
	PrivateExtension              []gtpType.PrivateExtension
	AdditionalIEs                 []gtpType.InformationElement
}

func NewErrorIndication(sequenceNumber uint16, teid uint32, peerAddress gtpType.GTPUPeerAddress) *ErrorIndication {
	h := gtpType.NewHeader(gtpType.MsgErrorIndication, 0)
	h.S = true
	h.SequenceNumber = sequenceNumber
	return &ErrorIndication{
		Header:                        h,
		TunnelEndpointIdentifierDataI: gtpType.NewTunnelEndpointIdentifierDataI(teid),
		GTPUPeerAddress:               peerAddress,
	}
}

func (m *ErrorIndication) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("Error Indication is nil")
	}
	if m.TunnelEndpointIdentifierDataI.Header.Type != gtpType.IETunnelEndpointIdentifierDataI {
		return nil, fmt.Errorf("mandatory Tunnel Endpoint Identifier Data I IE is missing")
	}
	if m.GTPUPeerAddress.Header.Type != gtpType.IEGTPUPeerAddress {
		return nil, fmt.Errorf("mandatory GTP-U Peer Address IE is missing")
	}

	h := m.Header
	h.Version = gtpType.Version1
	h.ProtocolType = true
	h.MessageType = gtpType.MsgErrorIndication
	h.TEID = 0
	h.S = true
	h.E = false
	h.PN = false
	h.NPDUNumber = 0
	h.NextExtensionHeaderType = 0

	type encodedIE struct {
		typeID uint8
		data   []byte
	}
	var ies []encodedIE

	b, err := m.TunnelEndpointIdentifierDataI.Encode()
	if err != nil {
		return nil, err
	}
	ies = append(ies, encodedIE{typeID: gtpType.IETunnelEndpointIdentifierDataI, data: b})

	b, err = m.GTPUPeerAddress.Encode()
	if err != nil {
		return nil, err
	}
	ies = append(ies, encodedIE{typeID: gtpType.IEGTPUPeerAddress, data: b})

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
	body := make([]byte, 0)
	for _, ie := range ies {
		body = append(body, ie.data...)
	}

	payloadLen := 4 + len(body) // S=1 => common optional block is present.
	if payloadLen > 0xffff {
		return nil, fmt.Errorf("Error Indication too large: %d", payloadLen)
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

func DecodeErrorIndication(data []byte) (*ErrorIndication, error) {
	h, err := gtpType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.MessageType != gtpType.MsgErrorIndication {
		return nil, fmt.Errorf("unexpected GTP message type: %d", h.MessageType)
	}
	if h.TEID != 0 {
		return nil, fmt.Errorf("Error Indication header TEID must be zero: 0x%08x", h.TEID)
	}
	if !h.S {
		return nil, fmt.Errorf("Error Indication S flag must be set")
	}
	if h.E || h.PN {
		return nil, fmt.Errorf("unsupported Error Indication header flags: E=%v PN=%v", h.E, h.PN)
	}

	m := &ErrorIndication{Header: h}
	offset := 12
	for offset < len(data) {
		ieType := data[offset]
		switch ieType {
		case gtpType.IETunnelEndpointIdentifierDataI:
			if m.TunnelEndpointIdentifierDataI.Header.Type != 0 {
				return nil, fmt.Errorf("duplicate Tunnel Endpoint Identifier Data I IE")
			}
			v, n, err := gtpType.DecodeTunnelEndpointIdentifierDataI(data[offset:])
			if err != nil {
				return nil, err
			}
			m.TunnelEndpointIdentifierDataI = v
			offset += n

		case gtpType.IEGTPUPeerAddress:
			if m.GTPUPeerAddress.Header.Type != 0 {
				return nil, fmt.Errorf("duplicate GTP-U Peer Address IE")
			}
			v, n, err := gtpType.DecodeGTPUPeerAddress(data[offset:])
			if err != nil {
				return nil, err
			}
			m.GTPUPeerAddress = v
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
				return nil, fmt.Errorf("unsupported TV IE in Error Indication: %d", ieType)
			}
			v, n, err := gtpType.DecodeTLVInformationElement(data[offset:])
			if err != nil {
				return nil, err
			}
			m.AdditionalIEs = append(m.AdditionalIEs, v)
			offset += n
		}
	}

	if m.TunnelEndpointIdentifierDataI.Header.Type != gtpType.IETunnelEndpointIdentifierDataI {
		return nil, fmt.Errorf("mandatory Tunnel Endpoint Identifier Data I IE is missing")
	}
	if m.GTPUPeerAddress.Header.Type != gtpType.IEGTPUPeerAddress {
		return nil, fmt.Errorf("mandatory GTP-U Peer Address IE is missing")
	}
	return m, nil
}
