// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EmergencyAreaIDBroadcastEUTRA represents ASN.1 EmergencyAreaIDBroadcastEUTRA.
type EmergencyAreaIDBroadcastEUTRA struct {
	List []EmergencyAreaIDBroadcastEUTRAItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyAreaIDBroadcastEUTRA) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyAreaIDBroadcastEUTRA", v)
}

// DecodeEmergencyAreaIDBroadcastEUTRA decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyAreaIDBroadcastEUTRA(data []byte) (EmergencyAreaIDBroadcastEUTRA, int, error) {
	var v EmergencyAreaIDBroadcastEUTRA
	err := ngapcodec.DecodeParameter("EmergencyAreaIDBroadcastEUTRA", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyAreaIDBroadcastEUTRA{}, 0, err
	}
	return v, len(data), err
}
