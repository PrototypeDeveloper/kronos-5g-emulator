// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ConfidentialityProtectionIndication represents ASN.1 ConfidentialityProtectionIndication.
const (
	ConfidentialityProtectionIndicationPresentRequired  Enumerated = 0
	ConfidentialityProtectionIndicationPresentPreferred Enumerated = 1
	ConfidentialityProtectionIndicationPresentNotNeeded Enumerated = 2
)

type ConfidentialityProtectionIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ConfidentialityProtectionIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ConfidentialityProtectionIndication", v)
}

// DecodeConfidentialityProtectionIndication decodes a standalone APER value and reports the consumed octets.
func DecodeConfidentialityProtectionIndication(data []byte) (ConfidentialityProtectionIndication, int, error) {
	var v ConfidentialityProtectionIndication
	err := ngapcodec.DecodeParameter("ConfidentialityProtectionIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return ConfidentialityProtectionIndication{}, 0, err
	}
	return v, len(data), err
}
