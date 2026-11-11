// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EmergencyAreaIDBroadcastNR represents ASN.1 EmergencyAreaIDBroadcastNR.
type EmergencyAreaIDBroadcastNR struct {
	List []EmergencyAreaIDBroadcastNRItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyAreaIDBroadcastNR) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyAreaIDBroadcastNR", v)
}

// DecodeEmergencyAreaIDBroadcastNR decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyAreaIDBroadcastNR(data []byte) (EmergencyAreaIDBroadcastNR, int, error) {
	var v EmergencyAreaIDBroadcastNR
	err := ngapcodec.DecodeParameter("EmergencyAreaIDBroadcastNR", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyAreaIDBroadcastNR{}, 0, err
	}
	return v, len(data), err
}
