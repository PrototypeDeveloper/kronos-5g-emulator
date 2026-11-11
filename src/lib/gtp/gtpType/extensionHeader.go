package gtpType

import "fmt"

const (
	ExtensionHeaderNone                uint8 = 0x00
	ExtensionHeaderNRRANContainer      uint8 = 0x84
	ExtensionHeaderPDUSessionContainer uint8 = 0x85
)

// ExtensionHeaderHeader is a local representation of extension-header metadata.
// Type is carried by the previous Next Extension Header Type field on the wire.
// Length is encoded in units of 4 octets and includes the length octet,
// content/padding, and the trailing Next Extension Header Type octet.
type ExtensionHeaderHeader struct {
	Type   uint8
	Length uint8
}

// ExtensionHeader preserves an extension header in wire order.
type ExtensionHeader struct {
	Header                  ExtensionHeaderHeader
	Content                 []byte
	NextExtensionHeaderType uint8
}

func NewExtensionHeader(extensionType uint8, content []byte) ExtensionHeader {
	c := append([]byte(nil), content...)
	return ExtensionHeader{
		Header:  ExtensionHeaderHeader{Type: extensionType},
		Content: c,
	}
}

func (e ExtensionHeader) Encode(nextType uint8) ([]byte, error) {
	if e.Header.Type == ExtensionHeaderNone {
		return nil, fmt.Errorf("extension header type is zero")
	}

	content := append([]byte(nil), e.Content...)
	total := 1 + len(content) + 1
	if rem := total % 4; rem != 0 {
		content = append(content, make([]byte, 4-rem)...)
		total = 1 + len(content) + 1
	}
	if total/4 > 255 {
		return nil, fmt.Errorf("extension header too long: %d", total)
	}

	b := make([]byte, total)
	b[0] = uint8(total / 4)
	copy(b[1:len(b)-1], content)
	b[len(b)-1] = nextType
	return b, nil
}

// DecodeExtensionHeaders decodes a chained list. startType is the type from
// the GTP optional field and offset points to the first extension header.
func DecodeExtensionHeaders(data []byte, offset int, startType uint8) ([]ExtensionHeader, int, error) {
	if startType == ExtensionHeaderNone {
		return nil, offset, nil
	}

	var headers []ExtensionHeader
	currentType := startType

	for currentType != ExtensionHeaderNone {
		if offset >= len(data) {
			return nil, offset, fmt.Errorf("extension header is truncated")
		}

		units := data[offset]
		if units == 0 {
			return nil, offset, fmt.Errorf("invalid extension header length: 0")
		}
		total := int(units) * 4
		if offset+total > len(data) {
			return nil, offset, fmt.Errorf("extension header length exceeds packet: type=0x%02x", currentType)
		}
		if total < 4 {
			return nil, offset, fmt.Errorf("invalid extension header length: %d", total)
		}

		nextType := data[offset+total-1]
		content := append([]byte(nil), data[offset+1:offset+total-1]...)
		headers = append(headers, ExtensionHeader{
			Header: ExtensionHeaderHeader{
				Type:   currentType,
				Length: units,
			},
			Content:                 content,
			NextExtensionHeaderType: nextType,
		})

		offset += total
		currentType = nextType
	}

	return headers, offset, nil
}
