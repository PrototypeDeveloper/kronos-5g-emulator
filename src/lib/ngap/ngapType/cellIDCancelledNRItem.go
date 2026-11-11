// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellIDCancelledNRItem represents ASN.1 CellIDCancelledNR-Item.
type CellIDCancelledNRItem struct {
	NRCGI              NRCGI
	NumberOfBroadcasts NumberOfBroadcasts
	HasIEExtensions    bool
	IEExtensions       []CellIDCancelledNRItemExtIEsEntry
	UnknownExtensions  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellIDCancelledNRItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellIDCancelledNR-Item", v)
}

// DecodeCellIDCancelledNRItem decodes a standalone APER value and reports the consumed octets.
func DecodeCellIDCancelledNRItem(data []byte) (CellIDCancelledNRItem, int, error) {
	var v CellIDCancelledNRItem
	err := ngapcodec.DecodeParameter("CellIDCancelledNR-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellIDCancelledNRItem{}, 0, err
	}
	return v, len(data), err
}
