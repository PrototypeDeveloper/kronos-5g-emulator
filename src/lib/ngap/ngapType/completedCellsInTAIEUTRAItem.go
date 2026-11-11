// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CompletedCellsInTAIEUTRAItem represents ASN.1 CompletedCellsInTAI-EUTRA-Item.
type CompletedCellsInTAIEUTRAItem struct {
	EUTRACGI          EUTRACGI
	HasIEExtensions   bool
	IEExtensions      []CompletedCellsInTAIEUTRAItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CompletedCellsInTAIEUTRAItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CompletedCellsInTAI-EUTRA-Item", v)
}

// DecodeCompletedCellsInTAIEUTRAItem decodes a standalone APER value and reports the consumed octets.
func DecodeCompletedCellsInTAIEUTRAItem(data []byte) (CompletedCellsInTAIEUTRAItem, int, error) {
	var v CompletedCellsInTAIEUTRAItem
	err := ngapcodec.DecodeParameter("CompletedCellsInTAI-EUTRA-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return CompletedCellsInTAIEUTRAItem{}, 0, err
	}
	return v, len(data), err
}
