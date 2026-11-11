// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CompletedCellsInTAINRItem represents ASN.1 CompletedCellsInTAI-NR-Item.
type CompletedCellsInTAINRItem struct {
	NRCGI             NRCGI
	HasIEExtensions   bool
	IEExtensions      []CompletedCellsInTAINRItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CompletedCellsInTAINRItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CompletedCellsInTAI-NR-Item", v)
}

// DecodeCompletedCellsInTAINRItem decodes a standalone APER value and reports the consumed octets.
func DecodeCompletedCellsInTAINRItem(data []byte) (CompletedCellsInTAINRItem, int, error) {
	var v CompletedCellsInTAINRItem
	err := ngapcodec.DecodeParameter("CompletedCellsInTAI-NR-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return CompletedCellsInTAINRItem{}, 0, err
	}
	return v, len(data), err
}
