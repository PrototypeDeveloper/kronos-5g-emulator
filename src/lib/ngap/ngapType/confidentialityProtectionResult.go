// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ConfidentialityProtectionResult represents ASN.1 ConfidentialityProtectionResult.
const (
	ConfidentialityProtectionResultPresentPerformed    Enumerated = 0
	ConfidentialityProtectionResultPresentNotPerformed Enumerated = 1
)

type ConfidentialityProtectionResult struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ConfidentialityProtectionResult) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ConfidentialityProtectionResult", v)
}

// DecodeConfidentialityProtectionResult decodes a standalone APER value and reports the consumed octets.
func DecodeConfidentialityProtectionResult(data []byte) (ConfidentialityProtectionResult, int, error) {
	var v ConfidentialityProtectionResult
	err := ngapcodec.DecodeParameter("ConfidentialityProtectionResult", data, &v)
	if ngapcodec.IsFatal(err) {
		return ConfidentialityProtectionResult{}, 0, err
	}
	return v, len(data), err
}
