// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellIDBroadcastNRItem represents ASN.1 CellIDBroadcastNR-Item.
type CellIDBroadcastNRItem struct {
	NRCGI             NRCGI
	HasIEExtensions   bool
	IEExtensions      []CellIDBroadcastNRItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellIDBroadcastNRItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellIDBroadcastNR-Item", v)
}

// DecodeCellIDBroadcastNRItem decodes a standalone APER value and reports the consumed octets.
func DecodeCellIDBroadcastNRItem(data []byte) (CellIDBroadcastNRItem, int, error) {
	var v CellIDBroadcastNRItem
	err := ngapcodec.DecodeParameter("CellIDBroadcastNR-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellIDBroadcastNRItem{}, 0, err
	}
	return v, len(data), err
}
