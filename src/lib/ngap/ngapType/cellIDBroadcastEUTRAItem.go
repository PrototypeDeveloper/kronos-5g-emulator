// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellIDBroadcastEUTRAItem represents ASN.1 CellIDBroadcastEUTRA-Item.
type CellIDBroadcastEUTRAItem struct {
	EUTRACGI          EUTRACGI
	HasIEExtensions   bool
	IEExtensions      []CellIDBroadcastEUTRAItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellIDBroadcastEUTRAItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellIDBroadcastEUTRA-Item", v)
}

// DecodeCellIDBroadcastEUTRAItem decodes a standalone APER value and reports the consumed octets.
func DecodeCellIDBroadcastEUTRAItem(data []byte) (CellIDBroadcastEUTRAItem, int, error) {
	var v CellIDBroadcastEUTRAItem
	err := ngapcodec.DecodeParameter("CellIDBroadcastEUTRA-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellIDBroadcastEUTRAItem{}, 0, err
	}
	return v, len(data), err
}
