// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BroadcastCompletedAreaList represents ASN.1 BroadcastCompletedAreaList.
type BroadcastCompletedAreaListPresent uint32

const (
	BroadcastCompletedAreaListPresentNothing BroadcastCompletedAreaListPresent = iota
	BroadcastCompletedAreaListPresentCellIDBroadcastEUTRA
	BroadcastCompletedAreaListPresentTAIBroadcastEUTRA
	BroadcastCompletedAreaListPresentEmergencyAreaIDBroadcastEUTRA
	BroadcastCompletedAreaListPresentCellIDBroadcastNR
	BroadcastCompletedAreaListPresentTAIBroadcastNR
	BroadcastCompletedAreaListPresentEmergencyAreaIDBroadcastNR
	BroadcastCompletedAreaListPresentChoiceExtensions
)

type BroadcastCompletedAreaList struct {
	Present                       BroadcastCompletedAreaListPresent
	CellIDBroadcastEUTRA          CellIDBroadcastEUTRA
	TAIBroadcastEUTRA             TAIBroadcastEUTRA
	EmergencyAreaIDBroadcastEUTRA EmergencyAreaIDBroadcastEUTRA
	CellIDBroadcastNR             CellIDBroadcastNR
	TAIBroadcastNR                TAIBroadcastNR
	EmergencyAreaIDBroadcastNR    EmergencyAreaIDBroadcastNR
	ChoiceExtensions              BroadcastCompletedAreaListExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BroadcastCompletedAreaList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BroadcastCompletedAreaList", v)
}

// DecodeBroadcastCompletedAreaList decodes a standalone APER value and reports the consumed octets.
func DecodeBroadcastCompletedAreaList(data []byte) (BroadcastCompletedAreaList, int, error) {
	var v BroadcastCompletedAreaList
	err := ngapcodec.DecodeParameter("BroadcastCompletedAreaList", data, &v)
	if ngapcodec.IsFatal(err) {
		return BroadcastCompletedAreaList{}, 0, err
	}
	return v, len(data), err
}
