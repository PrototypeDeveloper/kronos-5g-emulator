// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntegrityProtectionIndication represents ASN.1 IntegrityProtectionIndication.
const (
	IntegrityProtectionIndicationPresentRequired  Enumerated = 0
	IntegrityProtectionIndicationPresentPreferred Enumerated = 1
	IntegrityProtectionIndicationPresentNotNeeded Enumerated = 2
)

type IntegrityProtectionIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntegrityProtectionIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntegrityProtectionIndication", v)
}

// DecodeIntegrityProtectionIndication decodes a standalone APER value and reports the consumed octets.
func DecodeIntegrityProtectionIndication(data []byte) (IntegrityProtectionIndication, int, error) {
	var v IntegrityProtectionIndication
	err := ngapcodec.DecodeParameter("IntegrityProtectionIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntegrityProtectionIndication{}, 0, err
	}
	return v, len(data), err
}
