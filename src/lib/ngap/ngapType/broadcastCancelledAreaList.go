// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BroadcastCancelledAreaList represents ASN.1 BroadcastCancelledAreaList.
type BroadcastCancelledAreaListPresent uint32

const (
	BroadcastCancelledAreaListPresentNothing BroadcastCancelledAreaListPresent = iota
	BroadcastCancelledAreaListPresentCellIDCancelledEUTRA
	BroadcastCancelledAreaListPresentTAICancelledEUTRA
	BroadcastCancelledAreaListPresentEmergencyAreaIDCancelledEUTRA
	BroadcastCancelledAreaListPresentCellIDCancelledNR
	BroadcastCancelledAreaListPresentTAICancelledNR
	BroadcastCancelledAreaListPresentEmergencyAreaIDCancelledNR
	BroadcastCancelledAreaListPresentChoiceExtensions
)

type BroadcastCancelledAreaList struct {
	Present                       BroadcastCancelledAreaListPresent
	CellIDCancelledEUTRA          CellIDCancelledEUTRA
	TAICancelledEUTRA             TAICancelledEUTRA
	EmergencyAreaIDCancelledEUTRA EmergencyAreaIDCancelledEUTRA
	CellIDCancelledNR             CellIDCancelledNR
	TAICancelledNR                TAICancelledNR
	EmergencyAreaIDCancelledNR    EmergencyAreaIDCancelledNR
	ChoiceExtensions              BroadcastCancelledAreaListExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BroadcastCancelledAreaList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BroadcastCancelledAreaList", v)
}

// DecodeBroadcastCancelledAreaList decodes a standalone APER value and reports the consumed octets.
func DecodeBroadcastCancelledAreaList(data []byte) (BroadcastCancelledAreaList, int, error) {
	var v BroadcastCancelledAreaList
	err := ngapcodec.DecodeParameter("BroadcastCancelledAreaList", data, &v)
	if ngapcodec.IsFatal(err) {
		return BroadcastCancelledAreaList{}, 0, err
	}
	return v, len(data), err
}
