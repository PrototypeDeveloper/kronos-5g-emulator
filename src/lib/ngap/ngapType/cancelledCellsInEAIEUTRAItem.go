// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CancelledCellsInEAIEUTRAItem represents ASN.1 CancelledCellsInEAI-EUTRA-Item.
type CancelledCellsInEAIEUTRAItem struct {
	EUTRACGI           EUTRACGI
	NumberOfBroadcasts NumberOfBroadcasts
	HasIEExtensions    bool
	IEExtensions       []CancelledCellsInEAIEUTRAItemExtIEsEntry
	UnknownExtensions  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CancelledCellsInEAIEUTRAItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CancelledCellsInEAI-EUTRA-Item", v)
}

// DecodeCancelledCellsInEAIEUTRAItem decodes a standalone APER value and reports the consumed octets.
func DecodeCancelledCellsInEAIEUTRAItem(data []byte) (CancelledCellsInEAIEUTRAItem, int, error) {
	var v CancelledCellsInEAIEUTRAItem
	err := ngapcodec.DecodeParameter("CancelledCellsInEAI-EUTRA-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return CancelledCellsInEAIEUTRAItem{}, 0, err
	}
	return v, len(data), err
}
