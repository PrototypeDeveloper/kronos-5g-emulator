// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CompletedCellsInEAIEUTRAItem represents ASN.1 CompletedCellsInEAI-EUTRA-Item.
type CompletedCellsInEAIEUTRAItem struct {
	EUTRACGI          EUTRACGI
	HasIEExtensions   bool
	IEExtensions      []CompletedCellsInEAIEUTRAItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CompletedCellsInEAIEUTRAItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CompletedCellsInEAI-EUTRA-Item", v)
}

// DecodeCompletedCellsInEAIEUTRAItem decodes a standalone APER value and reports the consumed octets.
func DecodeCompletedCellsInEAIEUTRAItem(data []byte) (CompletedCellsInEAIEUTRAItem, int, error) {
	var v CompletedCellsInEAIEUTRAItem
	err := ngapcodec.DecodeParameter("CompletedCellsInEAI-EUTRA-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return CompletedCellsInEAIEUTRAItem{}, 0, err
	}
	return v, len(data), err
}
