// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PrivacyIndicator represents ASN.1 PrivacyIndicator.
const (
	PrivacyIndicatorPresentImmediateMDT Enumerated = 0
	PrivacyIndicatorPresentLoggedMDT    Enumerated = 1
)

type PrivacyIndicator struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PrivacyIndicator) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PrivacyIndicator", v)
}

// DecodePrivacyIndicator decodes a standalone APER value and reports the consumed octets.
func DecodePrivacyIndicator(data []byte) (PrivacyIndicator, int, error) {
	var v PrivacyIndicator
	err := ngapcodec.DecodeParameter("PrivacyIndicator", data, &v)
	if ngapcodec.IsFatal(err) {
		return PrivacyIndicator{}, 0, err
	}
	return v, len(data), err
}
