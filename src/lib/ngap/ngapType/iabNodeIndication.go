// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IABNodeIndication represents ASN.1 IABNodeIndication.
const (
	IABNodeIndicationPresentTrue Enumerated = 0
)

type IABNodeIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IABNodeIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IABNodeIndication", v)
}

// DecodeIABNodeIndication decodes a standalone APER value and reports the consumed octets.
func DecodeIABNodeIndication(data []byte) (IABNodeIndication, int, error) {
	var v IABNodeIndication
	err := ngapcodec.DecodeParameter("IABNodeIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return IABNodeIndication{}, 0, err
	}
	return v, len(data), err
}
