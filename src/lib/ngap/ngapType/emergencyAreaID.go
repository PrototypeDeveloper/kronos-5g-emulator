// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EmergencyAreaID represents ASN.1 EmergencyAreaID.
type EmergencyAreaID struct {
	Value []byte `aper:"sizeLB:3,sizeUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyAreaID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyAreaID", v)
}

// DecodeEmergencyAreaID decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyAreaID(data []byte) (EmergencyAreaID, int, error) {
	var v EmergencyAreaID
	err := ngapcodec.DecodeParameter("EmergencyAreaID", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyAreaID{}, 0, err
	}
	return v, len(data), err
}
