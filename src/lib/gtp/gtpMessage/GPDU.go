package gtpMessage

import (
	"fmt"

	"libgtp/gtpType"
)

// GPDU carries the original user packet/frame in TPDU.
// ExtensionHeaders preserves extension-header wire order. PDUSessionContainer
// is also exposed as a typed direct-access view for 5G N3/N9 processing.
type GPDU struct {
	Header gtpType.Header

	PDUSessionContainer gtpType.PDUSessionContainer
	ExtensionHeaders    []gtpType.ExtensionHeader

	TPDU []byte
}

func NewGPDU(teid uint32, tpdu []byte) *GPDU {
	return &GPDU{
		Header: gtpType.NewHeader(gtpType.MsgGPDU, teid),
		TPDU:   append([]byte(nil), tpdu...),
	}
}

func NewGPDUWithPDUSessionContainer(teid uint32, tpdu []byte, pduType, qfi uint8) (*GPDU, error) {
	psc, err := gtpType.NewPDUSessionContainer(pduType, qfi)
	if err != nil {
		return nil, err
	}
	m := NewGPDU(teid, tpdu)
	m.PDUSessionContainer = psc
	return m, nil
}

func (m *GPDU) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("G-PDU is nil")
	}

	h := m.Header
	h.Version = gtpType.Version1
	h.ProtocolType = true
	h.MessageType = gtpType.MsgGPDU

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

	payloadLength := optionalLength + extensionLength + len(m.TPDU)
	if payloadLength > 0xffff {
		return nil, fmt.Errorf("G-PDU payload too large: %d", payloadLength)
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
	out = append(out, m.TPDU...)

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

func DecodeGPDU(data []byte) (*GPDU, error) {
	h, err := gtpType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.MessageType != gtpType.MsgGPDU {
		return nil, fmt.Errorf("unexpected GTP message type: %d", h.MessageType)
	}

	m := &GPDU{Header: h}
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

	if offset > len(data) {
		return nil, fmt.Errorf("invalid G-PDU payload offset: %d", offset)
	}
	m.TPDU = append([]byte(nil), data[offset:]...)
	return m, nil
}

func cloneExtensionHeaders(in []gtpType.ExtensionHeader) []gtpType.ExtensionHeader {
	out := make([]gtpType.ExtensionHeader, len(in))
	for i := range in {
		out[i] = in[i]
		out[i].Content = append([]byte(nil), in[i].Content...)
	}
	return out
}
