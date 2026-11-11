// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EmergencyAreaIDCancelledNR represents ASN.1 EmergencyAreaIDCancelledNR.
type EmergencyAreaIDCancelledNR struct {
	List []EmergencyAreaIDCancelledNRItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyAreaIDCancelledNR) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyAreaIDCancelledNR", v)
}

// DecodeEmergencyAreaIDCancelledNR decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyAreaIDCancelledNR(data []byte) (EmergencyAreaIDCancelledNR, int, error) {
	var v EmergencyAreaIDCancelledNR
	err := ngapcodec.DecodeParameter("EmergencyAreaIDCancelledNR", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyAreaIDCancelledNR{}, 0, err
	}
	return v, len(data), err
}
