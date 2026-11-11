// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RedCapIndication represents ASN.1 RedCapIndication.
const (
	RedCapIndicationPresentRedcap Enumerated = 0
)

type RedCapIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RedCapIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RedCapIndication", v)
}

// DecodeRedCapIndication decodes a standalone APER value and reports the consumed octets.
func DecodeRedCapIndication(data []byte) (RedCapIndication, int, error) {
	var v RedCapIndication
	err := ngapcodec.DecodeParameter("RedCapIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return RedCapIndication{}, 0, err
	}
	return v, len(data), err
}
