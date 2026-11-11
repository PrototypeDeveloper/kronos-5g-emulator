// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EmergencyAreaIDCancelledEUTRAItem represents ASN.1 EmergencyAreaIDCancelledEUTRA-Item.
type EmergencyAreaIDCancelledEUTRAItem struct {
	EmergencyAreaID          EmergencyAreaID
	CancelledCellsInEAIEUTRA CancelledCellsInEAIEUTRA
	HasIEExtensions          bool
	IEExtensions             []EmergencyAreaIDCancelledEUTRAItemExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyAreaIDCancelledEUTRAItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyAreaIDCancelledEUTRA-Item", v)
}

// DecodeEmergencyAreaIDCancelledEUTRAItem decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyAreaIDCancelledEUTRAItem(data []byte) (EmergencyAreaIDCancelledEUTRAItem, int, error) {
	var v EmergencyAreaIDCancelledEUTRAItem
	err := ngapcodec.DecodeParameter("EmergencyAreaIDCancelledEUTRA-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyAreaIDCancelledEUTRAItem{}, 0, err
	}
	return v, len(data), err
}
