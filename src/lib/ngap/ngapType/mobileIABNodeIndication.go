// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MobileIABNodeIndication represents ASN.1 MobileIABNodeIndication.
const (
	MobileIABNodeIndicationPresentTrue Enumerated = 0
)

type MobileIABNodeIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MobileIABNodeIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MobileIABNodeIndication", v)
}

// DecodeMobileIABNodeIndication decodes a standalone APER value and reports the consumed octets.
func DecodeMobileIABNodeIndication(data []byte) (MobileIABNodeIndication, int, error) {
	var v MobileIABNodeIndication
	err := ngapcodec.DecodeParameter("MobileIABNodeIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return MobileIABNodeIndication{}, 0, err
	}
	return v, len(data), err
}
