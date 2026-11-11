// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CancelledCellsInEAINRItem represents ASN.1 CancelledCellsInEAI-NR-Item.
type CancelledCellsInEAINRItem struct {
	NRCGI              NRCGI
	NumberOfBroadcasts NumberOfBroadcasts
	HasIEExtensions    bool
	IEExtensions       []CancelledCellsInEAINRItemExtIEsEntry
	UnknownExtensions  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CancelledCellsInEAINRItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CancelledCellsInEAI-NR-Item", v)
}

// DecodeCancelledCellsInEAINRItem decodes a standalone APER value and reports the consumed octets.
func DecodeCancelledCellsInEAINRItem(data []byte) (CancelledCellsInEAINRItem, int, error) {
	var v CancelledCellsInEAINRItem
	err := ngapcodec.DecodeParameter("CancelledCellsInEAI-NR-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return CancelledCellsInEAINRItem{}, 0, err
	}
	return v, len(data), err
}
