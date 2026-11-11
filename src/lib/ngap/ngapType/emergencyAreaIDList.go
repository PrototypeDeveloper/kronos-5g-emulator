// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EmergencyAreaIDList represents ASN.1 EmergencyAreaIDList.
type EmergencyAreaIDList struct {
	List []EmergencyAreaID `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyAreaIDList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyAreaIDList", v)
}

// DecodeEmergencyAreaIDList decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyAreaIDList(data []byte) (EmergencyAreaIDList, int, error) {
	var v EmergencyAreaIDList
	err := ngapcodec.DecodeParameter("EmergencyAreaIDList", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyAreaIDList{}, 0, err
	}
	return v, len(data), err
}
