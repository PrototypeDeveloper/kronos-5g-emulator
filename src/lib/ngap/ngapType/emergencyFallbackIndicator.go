// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EmergencyFallbackIndicator represents ASN.1 EmergencyFallbackIndicator.
type EmergencyFallbackIndicator struct {
	EmergencyFallbackRequestIndicator EmergencyFallbackRequestIndicator
	HasEmergencyServiceTargetCN       bool
	EmergencyServiceTargetCN          EmergencyServiceTargetCN
	HasIEExtensions                   bool
	IEExtensions                      []EmergencyFallbackIndicatorExtIEsEntry
	UnknownExtensions                 []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyFallbackIndicator) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyFallbackIndicator", v)
}

// DecodeEmergencyFallbackIndicator decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyFallbackIndicator(data []byte) (EmergencyFallbackIndicator, int, error) {
	var v EmergencyFallbackIndicator
	err := ngapcodec.DecodeParameter("EmergencyFallbackIndicator", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyFallbackIndicator{}, 0, err
	}
	return v, len(data), err
}
