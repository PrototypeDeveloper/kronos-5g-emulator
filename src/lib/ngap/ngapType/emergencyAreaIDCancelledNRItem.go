// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EmergencyAreaIDCancelledNRItem represents ASN.1 EmergencyAreaIDCancelledNR-Item.
type EmergencyAreaIDCancelledNRItem struct {
	EmergencyAreaID       EmergencyAreaID
	CancelledCellsInEAINR CancelledCellsInEAINR
	HasIEExtensions       bool
	IEExtensions          []EmergencyAreaIDCancelledNRItemExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyAreaIDCancelledNRItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyAreaIDCancelledNR-Item", v)
}

// DecodeEmergencyAreaIDCancelledNRItem decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyAreaIDCancelledNRItem(data []byte) (EmergencyAreaIDCancelledNRItem, int, error) {
	var v EmergencyAreaIDCancelledNRItem
	err := ngapcodec.DecodeParameter("EmergencyAreaIDCancelledNR-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyAreaIDCancelledNRItem{}, 0, err
	}
	return v, len(data), err
}
