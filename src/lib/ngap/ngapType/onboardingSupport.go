// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// OnboardingSupport represents ASN.1 OnboardingSupport.
const (
	OnboardingSupportPresentTrue Enumerated = 0
)

type OnboardingSupport struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v OnboardingSupport) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("OnboardingSupport", v)
}

// DecodeOnboardingSupport decodes a standalone APER value and reports the consumed octets.
func DecodeOnboardingSupport(data []byte) (OnboardingSupport, int, error) {
	var v OnboardingSupport
	err := ngapcodec.DecodeParameter("OnboardingSupport", data, &v)
	if ngapcodec.IsFatal(err) {
		return OnboardingSupport{}, 0, err
	}
	return v, len(data), err
}
