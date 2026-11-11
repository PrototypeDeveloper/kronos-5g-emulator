package gtpMessage

import (
	"fmt"
	"sort"

	"libgtp/gtpType"
)

// EndMarker indicates the end of the payload stream on a GTP-U tunnel.
// The tunnel is identified by Header.TEID. In 5GS data-forwarding scenarios,
// a PDU Session Container extension header may be present so the QFI can be
// accessed directly by the caller.
type EndMarker struct {
	Header gtpType.Header

	PDUSessionContainer gtpType.PDUSessionContainer
	ExtensionHeaders    []gtpType.ExtensionHeader

	PrivateExtension []gtpType.PrivateExtension
	AdditionalIEs    []gtpType.InformationElement
}

func NewEndMarker(teid uint32) *EndMarker {
	return &EndMarker{Header: gtpType.NewHeader(gtpType.MsgEndMarker, teid)}
}

func NewEndMarkerWithPDUSessionContainer(teid uint32, pduType, qfi uint8) (*EndMarker, error) {
	psc, err := gtpType.NewPDUSessionContainer(pduType, qfi)
	if err != nil {
		return nil, err
	}
	m := NewEndMarker(teid)
	m.PDUSessionContainer = psc
	return m, nil
}

func (m *EndMarker) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("End Marker is nil")
	}
	if m.Header.TEID == 0 {
		return nil, fmt.Errorf("End Marker TEID must not be zero")
	}

	h := m.Header
	h.Version = gtpType.Version1
	h.ProtocolType = true
	h.MessageType = gtpType.MsgEndMarker

	extHeaders := cloneExtensionHeaders(m.ExtensionHeaders)
	if m.PDUSessionContainer.IsPresent() {
		content, err := m.PDUSessionContainer.EncodeContent()
		if err != nil {
			return nil, err
		}
		found := false
		for i := range extHeaders {
			if extHeaders[i].Header.Type == gtpType.ExtensionHeaderPDUSessionContainer {
				extHeaders[i].Content = content
				found = true
				break
			}
		}
		if !found {
			extHeaders = append(extHeaders, gtpType.NewExtensionHeader(
				gtpType.ExtensionHeaderPDUSessionContainer,
				content,
			))
		}
	}

	h.E = len(extHeaders) > 0
	if h.E {
		h.NextExtensionHeaderType = extHeaders[0].Header.Type
	} else {
		h.NextExtensionHeaderType = 0
	}

	optionalLength := 0
	if h.HasOptionalFields() {
		optionalLength = 4
	}

	encodedExt := make([][]byte, 0, len(extHeaders))
	extensionLength := 0
	for i := range extHeaders {
		nextType := uint8(0)
		if i+1 < len(extHeaders) {
			nextType = extHeaders[i+1].Header.Type
		}
		b, err := extHeaders[i].Encode(nextType)
		if err != nil {
			return nil, err
		}
		encodedExt = append(encodedExt, b)
		extensionLength += len(b)
	}

	type encodedIE struct {
		typeID uint8
		data   []byte
	}
	ies := make([]encodedIE, 0, len(m.PrivateExtension)+len(m.AdditionalIEs))
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

	payloadLength := optionalLength + extensionLength + len(body)
	if payloadLength > 0xffff {
		return nil, fmt.Errorf("End Marker too large: %d", payloadLength)
	}
	h.Length = uint16(payloadLength)

	mandatory, err := h.EncodeMandatory()
	if err != nil {
		return nil, err
	}

	out := make([]byte, 0, 8+payloadLength)
	out = append(out, mandatory...)
	out = append(out, h.EncodeOptional()...)
	for _, b := range encodedExt {
		out = append(out, b...)
	}
	out = append(out, body...)

	m.Header = h
	m.ExtensionHeaders = extHeaders
	for i := range m.ExtensionHeaders {
		if i+1 < len(m.ExtensionHeaders) {
			m.ExtensionHeaders[i].NextExtensionHeaderType = m.ExtensionHeaders[i+1].Header.Type
		} else {
			m.ExtensionHeaders[i].NextExtensionHeaderType = 0
		}
		m.ExtensionHeaders[i].Header.Length = uint8(len(encodedExt[i]) / 4)
	}
	if m.PDUSessionContainer.IsPresent() {
		for _, e := range m.ExtensionHeaders {
			if e.Header.Type == gtpType.ExtensionHeaderPDUSessionContainer {
				m.PDUSessionContainer.Header.Length = e.Header.Length
				break
			}
		}
	}

	return out, nil
}

func DecodeEndMarker(data []byte) (*EndMarker, error) {
	h, err := gtpType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.MessageType != gtpType.MsgEndMarker {
		return nil, fmt.Errorf("unexpected GTP message type: %d", h.MessageType)
	}
	if h.TEID == 0 {
		return nil, fmt.Errorf("End Marker TEID must not be zero")
	}

	m := &EndMarker{Header: h}
	offset := 8
	if h.HasOptionalFields() {
		offset = 12
	}

	if h.E {
		if h.NextExtensionHeaderType == gtpType.ExtensionHeaderNone {
			return nil, fmt.Errorf("E flag is set but Next Extension Header Type is zero")
		}
		headers, nextOffset, err := gtpType.DecodeExtensionHeaders(data, offset, h.NextExtensionHeaderType)
		if err != nil {
			return nil, err
		}
		m.ExtensionHeaders = headers
		offset = nextOffset

		for _, e := range headers {
			if e.Header.Type == gtpType.ExtensionHeaderPDUSessionContainer {
				psc, err := gtpType.DecodePDUSessionContainer(e)
				if err != nil {
					return nil, err
				}
				m.PDUSessionContainer = psc
				break
			}
		}
	}

	for offset < len(data) {
		ieType := data[offset]
		switch ieType {
		case gtpType.IEPrivateExtension:
			v, n, err := gtpType.DecodePrivateExtension(data[offset:])
			if err != nil {
				return nil, err
			}
			m.PrivateExtension = append(m.PrivateExtension, v)
			offset += n

		default:
			if ieType < 128 {
				return nil, fmt.Errorf("unsupported TV IE in End Marker: %d", ieType)
			}
			v, n, err := gtpType.DecodeTLVInformationElement(data[offset:])
			if err != nil {
				return nil, err
			}
			m.AdditionalIEs = append(m.AdditionalIEs, v)
			offset += n
		}
	}

	return m, nil
}
