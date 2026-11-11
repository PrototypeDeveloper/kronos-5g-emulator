// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EmergencyAreaIDBroadcastNRItem represents ASN.1 EmergencyAreaIDBroadcastNR-Item.
type EmergencyAreaIDBroadcastNRItem struct {
	EmergencyAreaID       EmergencyAreaID
	CompletedCellsInEAINR CompletedCellsInEAINR
	HasIEExtensions       bool
	IEExtensions          []EmergencyAreaIDBroadcastNRItemExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyAreaIDBroadcastNRItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyAreaIDBroadcastNR-Item", v)
}

// DecodeEmergencyAreaIDBroadcastNRItem decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyAreaIDBroadcastNRItem(data []byte) (EmergencyAreaIDBroadcastNRItem, int, error) {
	var v EmergencyAreaIDBroadcastNRItem
	err := ngapcodec.DecodeParameter("EmergencyAreaIDBroadcastNR-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyAreaIDBroadcastNRItem{}, 0, err
	}
	return v, len(data), err
}
