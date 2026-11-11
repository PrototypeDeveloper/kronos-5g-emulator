// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GUAMIType represents ASN.1 GUAMIType.
const (
	GUAMITypePresentNative Enumerated = 0
	GUAMITypePresentMapped Enumerated = 1
)

type GUAMIType struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GUAMIType) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("GUAMIType", v) }

// DecodeGUAMIType decodes a standalone APER value and reports the consumed octets.
func DecodeGUAMIType(data []byte) (GUAMIType, int, error) {
	var v GUAMIType
	err := ngapcodec.DecodeParameter("GUAMIType", data, &v)
	if ngapcodec.IsFatal(err) {
		return GUAMIType{}, 0, err
	}
	return v, len(data), err
}
