// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CompletedCellsInEAINRItem represents ASN.1 CompletedCellsInEAI-NR-Item.
type CompletedCellsInEAINRItem struct {
	NRCGI             NRCGI
	HasIEExtensions   bool
	IEExtensions      []CompletedCellsInEAINRItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CompletedCellsInEAINRItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CompletedCellsInEAI-NR-Item", v)
}

// DecodeCompletedCellsInEAINRItem decodes a standalone APER value and reports the consumed octets.
func DecodeCompletedCellsInEAINRItem(data []byte) (CompletedCellsInEAINRItem, int, error) {
	var v CompletedCellsInEAINRItem
	err := ngapcodec.DecodeParameter("CompletedCellsInEAI-NR-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return CompletedCellsInEAINRItem{}, 0, err
	}
	return v, len(data), err
}
