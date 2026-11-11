// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CancelledCellsInTAINRItem represents ASN.1 CancelledCellsInTAI-NR-Item.
type CancelledCellsInTAINRItem struct {
	NRCGI              NRCGI
	NumberOfBroadcasts NumberOfBroadcasts
	HasIEExtensions    bool
	IEExtensions       []CancelledCellsInTAINRItemExtIEsEntry
	UnknownExtensions  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CancelledCellsInTAINRItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CancelledCellsInTAI-NR-Item", v)
}

// DecodeCancelledCellsInTAINRItem decodes a standalone APER value and reports the consumed octets.
func DecodeCancelledCellsInTAINRItem(data []byte) (CancelledCellsInTAINRItem, int, error) {
	var v CancelledCellsInTAINRItem
	err := ngapcodec.DecodeParameter("CancelledCellsInTAI-NR-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return CancelledCellsInTAINRItem{}, 0, err
	}
	return v, len(data), err
}
