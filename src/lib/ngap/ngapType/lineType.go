// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LineType represents ASN.1 LineType.
const (
	LineTypePresentDsl Enumerated = 0
	LineTypePresentPon Enumerated = 1
)

type LineType struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LineType) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("LineType", v) }

// DecodeLineType decodes a standalone APER value and reports the consumed octets.
func DecodeLineType(data []byte) (LineType, int, error) {
	var v LineType
	err := ngapcodec.DecodeParameter("LineType", data, &v)
	if ngapcodec.IsFatal(err) {
		return LineType{}, 0, err
	}
	return v, len(data), err
}
