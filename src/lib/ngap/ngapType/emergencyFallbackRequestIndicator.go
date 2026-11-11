// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EmergencyFallbackRequestIndicator represents ASN.1 EmergencyFallbackRequestIndicator.
const (
	EmergencyFallbackRequestIndicatorPresentEmergencyFallbackRequested Enumerated = 0
)

type EmergencyFallbackRequestIndicator struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyFallbackRequestIndicator) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyFallbackRequestIndicator", v)
}

// DecodeEmergencyFallbackRequestIndicator decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyFallbackRequestIndicator(data []byte) (EmergencyFallbackRequestIndicator, int, error) {
	var v EmergencyFallbackRequestIndicator
	err := ngapcodec.DecodeParameter("EmergencyFallbackRequestIndicator", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyFallbackRequestIndicator{}, 0, err
	}
	return v, len(data), err
}
