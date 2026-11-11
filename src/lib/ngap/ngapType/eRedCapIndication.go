// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ERedCapIndication represents ASN.1 ERedCapIndication.
const (
	ERedCapIndicationPresentTrue Enumerated = 0
)

type ERedCapIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ERedCapIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ERedCapIndication", v)
}

// DecodeERedCapIndication decodes a standalone APER value and reports the consumed octets.
func DecodeERedCapIndication(data []byte) (ERedCapIndication, int, error) {
	var v ERedCapIndication
	err := ngapcodec.DecodeParameter("ERedCapIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return ERedCapIndication{}, 0, err
	}
	return v, len(data), err
}
