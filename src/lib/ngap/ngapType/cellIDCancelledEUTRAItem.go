// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellIDCancelledEUTRAItem represents ASN.1 CellIDCancelledEUTRA-Item.
type CellIDCancelledEUTRAItem struct {
	EUTRACGI           EUTRACGI
	NumberOfBroadcasts NumberOfBroadcasts
	HasIEExtensions    bool
	IEExtensions       []CellIDCancelledEUTRAItemExtIEsEntry
	UnknownExtensions  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellIDCancelledEUTRAItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellIDCancelledEUTRA-Item", v)
}

// DecodeCellIDCancelledEUTRAItem decodes a standalone APER value and reports the consumed octets.
func DecodeCellIDCancelledEUTRAItem(data []byte) (CellIDCancelledEUTRAItem, int, error) {
	var v CellIDCancelledEUTRAItem
	err := ngapcodec.DecodeParameter("CellIDCancelledEUTRA-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellIDCancelledEUTRAItem{}, 0, err
	}
	return v, len(data), err
}
