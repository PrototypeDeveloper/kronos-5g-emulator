// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EmergencyServiceTargetCN represents ASN.1 EmergencyServiceTargetCN.
const (
	EmergencyServiceTargetCNPresentFiveGC Enumerated = 0
	EmergencyServiceTargetCNPresentEpc    Enumerated = 1
)

type EmergencyServiceTargetCN struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyServiceTargetCN) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyServiceTargetCN", v)
}

// DecodeEmergencyServiceTargetCN decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyServiceTargetCN(data []byte) (EmergencyServiceTargetCN, int, error) {
	var v EmergencyServiceTargetCN
	err := ngapcodec.DecodeParameter("EmergencyServiceTargetCN", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyServiceTargetCN{}, 0, err
	}
	return v, len(data), err
}
