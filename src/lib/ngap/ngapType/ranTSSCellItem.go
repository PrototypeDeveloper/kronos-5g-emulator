// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANTSSCellItem represents ASN.1 RANTSSCellItem.
type RANTSSCellItem struct {
	NRCGI             NRCGI
	HasIEExtensions   bool
	IEExtensions      []RANTSSCellItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANTSSCellItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RANTSSCellItem", v)
}

// DecodeRANTSSCellItem decodes a standalone APER value and reports the consumed octets.
func DecodeRANTSSCellItem(data []byte) (RANTSSCellItem, int, error) {
	var v RANTSSCellItem
	err := ngapcodec.DecodeParameter("RANTSSCellItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANTSSCellItem{}, 0, err
	}
	return v, len(data), err
}
