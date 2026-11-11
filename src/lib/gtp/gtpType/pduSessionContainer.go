package gtpType

import "fmt"

const (
	PDUSessionInformationDL uint8 = 0
	PDUSessionInformationUL uint8 = 1
)

// PDUSessionContainer represents the 5G PDU Session Container extension header.
// PDUType and QFI are decoded explicitly because they are the fields needed most
// frequently by the UPF. RawContent is retained so later TS 38.415 fields can be
// added without losing unknown/future information.
type PDUSessionContainer struct {
	Header     ExtensionHeaderHeader
	PDUType    uint8
	QFI        uint8
	RawContent []byte
}

func NewPDUSessionContainer(pduType, qfi uint8) (PDUSessionContainer, error) {
	var p PDUSessionContainer
	if pduType > 15 {
		return p, fmt.Errorf("invalid PDU type: %d", pduType)
	}
	if qfi > 63 {
		return p, fmt.Errorf("invalid QFI: %d", qfi)
	}

	p.Header.Type = ExtensionHeaderPDUSessionContainer
	p.PDUType = pduType
	p.QFI = qfi
	p.RawContent = []byte{pduType << 4, qfi & 0x3f}
	return p, nil
}

func DecodePDUSessionContainer(e ExtensionHeader) (PDUSessionContainer, error) {
	var p PDUSessionContainer
	if e.Header.Type != ExtensionHeaderPDUSessionContainer {
		return p, fmt.Errorf("unexpected extension header type: 0x%02x", e.Header.Type)
	}
	if len(e.Content) < 2 {
		return p, fmt.Errorf("PDU Session Container too short: %d", len(e.Content))
	}

	p.Header = e.Header
	p.PDUType = (e.Content[0] >> 4) & 0x0f
	p.QFI = e.Content[1] & 0x3f
	p.RawContent = append([]byte(nil), e.Content...)
	return p, nil
}

func (p PDUSessionContainer) EncodeContent() ([]byte, error) {
	if p.Header.Type != ExtensionHeaderPDUSessionContainer {
		return nil, fmt.Errorf("PDU Session Container is not present")
	}
	if p.PDUType > 15 {
		return nil, fmt.Errorf("invalid PDU type: %d", p.PDUType)
	}
	if p.QFI > 63 {
		return nil, fmt.Errorf("invalid QFI: %d", p.QFI)
	}

	content := append([]byte(nil), p.RawContent...)
	if len(content) < 2 {
		content = make([]byte, 2)
	}
	content[0] = (content[0] & 0x0f) | ((p.PDUType & 0x0f) << 4)
	content[1] = (content[1] & 0xc0) | (p.QFI & 0x3f)
	return content, nil
}

func (p PDUSessionContainer) IsPresent() bool {
	return p.Header.Type == ExtensionHeaderPDUSessionContainer
}
