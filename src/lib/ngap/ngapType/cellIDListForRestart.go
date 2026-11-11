// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellIDListForRestart represents ASN.1 CellIDListForRestart.
type CellIDListForRestartPresent uint32

const (
	CellIDListForRestartPresentNothing CellIDListForRestartPresent = iota
	CellIDListForRestartPresentEUTRACGIListforRestart
	CellIDListForRestartPresentNRCGIListforRestart
	CellIDListForRestartPresentChoiceExtensions
)

type CellIDListForRestart struct {
	Present                CellIDListForRestartPresent
	EUTRACGIListforRestart EUTRACGIList
	NRCGIListforRestart    NRCGIList
	ChoiceExtensions       CellIDListForRestartExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellIDListForRestart) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellIDListForRestart", v)
}

// DecodeCellIDListForRestart decodes a standalone APER value and reports the consumed octets.
func DecodeCellIDListForRestart(data []byte) (CellIDListForRestart, int, error) {
	var v CellIDListForRestart
	err := ngapcodec.DecodeParameter("CellIDListForRestart", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellIDListForRestart{}, 0, err
	}
	return v, len(data), err
}
