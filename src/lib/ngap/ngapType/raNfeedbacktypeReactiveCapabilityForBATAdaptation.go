// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANfeedbacktypeReactiveCapabilityForBATAdaptation represents ASN.1 RANfeedbacktypeReactiveCapabilityForBATAdaptation.
const (
	RANfeedbacktypeReactiveCapabilityForBATAdaptationPresentTrue Enumerated = 0
)

type RANfeedbacktypeReactiveCapabilityForBATAdaptation struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v RANfeedbacktypeReactiveCapabilityForBATAdaptation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:RANfeedbacktypeReactiveCapabilityForBATAdaptation", v)
}

// DecodeRANfeedbacktypeReactiveCapabilityForBATAdaptation decodes a standalone APER value and reports consumed octets.
func DecodeRANfeedbacktypeReactiveCapabilityForBATAdaptation(data []byte) (RANfeedbacktypeReactiveCapabilityForBATAdaptation, int, error) {
	var v RANfeedbacktypeReactiveCapabilityForBATAdaptation
	err := ngapcodec.DecodeParameter("inline:RANfeedbacktypeReactiveCapabilityForBATAdaptation", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANfeedbacktypeReactiveCapabilityForBATAdaptation{}, 0, err
	}
	return v, len(data), err
}
