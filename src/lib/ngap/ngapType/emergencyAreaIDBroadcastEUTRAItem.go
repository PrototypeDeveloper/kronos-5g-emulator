// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EmergencyAreaIDBroadcastEUTRAItem represents ASN.1 EmergencyAreaIDBroadcastEUTRA-Item.
type EmergencyAreaIDBroadcastEUTRAItem struct {
	EmergencyAreaID          EmergencyAreaID
	CompletedCellsInEAIEUTRA CompletedCellsInEAIEUTRA
	HasIEExtensions          bool
	IEExtensions             []EmergencyAreaIDBroadcastEUTRAItemExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyAreaIDBroadcastEUTRAItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyAreaIDBroadcastEUTRA-Item", v)
}

// DecodeEmergencyAreaIDBroadcastEUTRAItem decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyAreaIDBroadcastEUTRAItem(data []byte) (EmergencyAreaIDBroadcastEUTRAItem, int, error) {
	var v EmergencyAreaIDBroadcastEUTRAItem
	err := ngapcodec.DecodeParameter("EmergencyAreaIDBroadcastEUTRA-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyAreaIDBroadcastEUTRAItem{}, 0, err
	}
	return v, len(data), err
}
