// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EmergencyAreaIDListForRestart represents ASN.1 EmergencyAreaIDListForRestart.
type EmergencyAreaIDListForRestart struct {
	List []EmergencyAreaID `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyAreaIDListForRestart) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyAreaIDListForRestart", v)
}

// DecodeEmergencyAreaIDListForRestart decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyAreaIDListForRestart(data []byte) (EmergencyAreaIDListForRestart, int, error) {
	var v EmergencyAreaIDListForRestart
	err := ngapcodec.DecodeParameter("EmergencyAreaIDListForRestart", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyAreaIDListForRestart{}, 0, err
	}
	return v, len(data), err
}
