// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// WarningAreaList represents ASN.1 WarningAreaList.
type WarningAreaListPresent uint32

const (
	WarningAreaListPresentNothing WarningAreaListPresent = iota
	WarningAreaListPresentEUTRACGIListForWarning
	WarningAreaListPresentNRCGIListForWarning
	WarningAreaListPresentTAIListForWarning
	WarningAreaListPresentEmergencyAreaIDList
	WarningAreaListPresentChoiceExtensions
)

type WarningAreaList struct {
	Present                WarningAreaListPresent
	EUTRACGIListForWarning EUTRACGIListForWarning
	NRCGIListForWarning    NRCGIListForWarning
	TAIListForWarning      TAIListForWarning
	EmergencyAreaIDList    EmergencyAreaIDList
	ChoiceExtensions       WarningAreaListExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v WarningAreaList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("WarningAreaList", v)
}

// DecodeWarningAreaList decodes a standalone APER value and reports the consumed octets.
func DecodeWarningAreaList(data []byte) (WarningAreaList, int, error) {
	var v WarningAreaList
	err := ngapcodec.DecodeParameter("WarningAreaList", data, &v)
	if ngapcodec.IsFatal(err) {
		return WarningAreaList{}, 0, err
	}
	return v, len(data), err
}
