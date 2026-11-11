// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntegrityProtectionResult represents ASN.1 IntegrityProtectionResult.
const (
	IntegrityProtectionResultPresentPerformed    Enumerated = 0
	IntegrityProtectionResultPresentNotPerformed Enumerated = 1
)

type IntegrityProtectionResult struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntegrityProtectionResult) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntegrityProtectionResult", v)
}

// DecodeIntegrityProtectionResult decodes a standalone APER value and reports the consumed octets.
func DecodeIntegrityProtectionResult(data []byte) (IntegrityProtectionResult, int, error) {
	var v IntegrityProtectionResult
	err := ngapcodec.DecodeParameter("IntegrityProtectionResult", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntegrityProtectionResult{}, 0, err
	}
	return v, len(data), err
}
