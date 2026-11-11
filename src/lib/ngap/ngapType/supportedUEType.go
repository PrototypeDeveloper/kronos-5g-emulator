// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SupportedUEType represents ASN.1 SupportedUEType.
const (
	SupportedUETypePresentNonRedCapERedCapUE Enumerated = 0
	SupportedUETypePresentRedCapERedCapUE    Enumerated = 1
)

type SupportedUEType struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SupportedUEType) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SupportedUEType", v)
}

// DecodeSupportedUEType decodes a standalone APER value and reports the consumed octets.
func DecodeSupportedUEType(data []byte) (SupportedUEType, int, error) {
	var v SupportedUEType
	err := ngapcodec.DecodeParameter("SupportedUEType", data, &v)
	if ngapcodec.IsFatal(err) {
		return SupportedUEType{}, 0, err
	}
	return v, len(data), err
}
