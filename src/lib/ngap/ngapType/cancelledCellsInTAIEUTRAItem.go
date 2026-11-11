// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CancelledCellsInTAIEUTRAItem represents ASN.1 CancelledCellsInTAI-EUTRA-Item.
type CancelledCellsInTAIEUTRAItem struct {
	EUTRACGI           EUTRACGI
	NumberOfBroadcasts NumberOfBroadcasts
	HasIEExtensions    bool
	IEExtensions       []CancelledCellsInTAIEUTRAItemExtIEsEntry
	UnknownExtensions  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CancelledCellsInTAIEUTRAItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CancelledCellsInTAI-EUTRA-Item", v)
}

// DecodeCancelledCellsInTAIEUTRAItem decodes a standalone APER value and reports the consumed octets.
func DecodeCancelledCellsInTAIEUTRAItem(data []byte) (CancelledCellsInTAIEUTRAItem, int, error) {
	var v CancelledCellsInTAIEUTRAItem
	err := ngapcodec.DecodeParameter("CancelledCellsInTAI-EUTRA-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return CancelledCellsInTAIEUTRAItem{}, 0, err
	}
	return v, len(data), err
}
