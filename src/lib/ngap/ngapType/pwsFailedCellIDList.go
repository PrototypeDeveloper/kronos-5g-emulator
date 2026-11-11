// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PWSFailedCellIDList represents ASN.1 PWSFailedCellIDList.
type PWSFailedCellIDListPresent uint32

const (
	PWSFailedCellIDListPresentNothing PWSFailedCellIDListPresent = iota
	PWSFailedCellIDListPresentEUTRACGIPWSFailedList
	PWSFailedCellIDListPresentNRCGIPWSFailedList
	PWSFailedCellIDListPresentChoiceExtensions
)

type PWSFailedCellIDList struct {
	Present               PWSFailedCellIDListPresent
	EUTRACGIPWSFailedList EUTRACGIList
	NRCGIPWSFailedList    NRCGIList
	ChoiceExtensions      PWSFailedCellIDListExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PWSFailedCellIDList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PWSFailedCellIDList", v)
}

// DecodePWSFailedCellIDList decodes a standalone APER value and reports the consumed octets.
func DecodePWSFailedCellIDList(data []byte) (PWSFailedCellIDList, int, error) {
	var v PWSFailedCellIDList
	err := ngapcodec.DecodeParameter("PWSFailedCellIDList", data, &v)
	if ngapcodec.IsFatal(err) {
		return PWSFailedCellIDList{}, 0, err
	}
	return v, len(data), err
}
