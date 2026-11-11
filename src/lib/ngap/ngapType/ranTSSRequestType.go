// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANTSSRequestType represents ASN.1 RAN-TSSRequestType.
const (
	RANTSSRequestTypePresentStart Enumerated = 0
	RANTSSRequestTypePresentStop  Enumerated = 1
)

type RANTSSRequestType struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANTSSRequestType) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RAN-TSSRequestType", v)
}

// DecodeRANTSSRequestType decodes a standalone APER value and reports the consumed octets.
func DecodeRANTSSRequestType(data []byte) (RANTSSRequestType, int, error) {
	var v RANTSSRequestType
	err := ngapcodec.DecodeParameter("RAN-TSSRequestType", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANTSSRequestType{}, 0, err
	}
	return v, len(data), err
}
