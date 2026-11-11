// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EmergencyAreaIDCancelledEUTRA represents ASN.1 EmergencyAreaIDCancelledEUTRA.
type EmergencyAreaIDCancelledEUTRA struct {
	List []EmergencyAreaIDCancelledEUTRAItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyAreaIDCancelledEUTRA) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyAreaIDCancelledEUTRA", v)
}

// DecodeEmergencyAreaIDCancelledEUTRA decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyAreaIDCancelledEUTRA(data []byte) (EmergencyAreaIDCancelledEUTRA, int, error) {
	var v EmergencyAreaIDCancelledEUTRA
	err := ngapcodec.DecodeParameter("EmergencyAreaIDCancelledEUTRA", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyAreaIDCancelledEUTRA{}, 0, err
	}
	return v, len(data), err
}
