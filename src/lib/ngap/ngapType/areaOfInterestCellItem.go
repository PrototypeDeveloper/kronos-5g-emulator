// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AreaOfInterestCellItem represents ASN.1 AreaOfInterestCellItem.
type AreaOfInterestCellItem struct {
	NGRANCGI          NGRANCGI
	HasIEExtensions   bool
	IEExtensions      []AreaOfInterestCellItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaOfInterestCellItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaOfInterestCellItem", v)
}

// DecodeAreaOfInterestCellItem decodes a standalone APER value and reports the consumed octets.
func DecodeAreaOfInterestCellItem(data []byte) (AreaOfInterestCellItem, int, error) {
	var v AreaOfInterestCellItem
	err := ngapcodec.DecodeParameter("AreaOfInterestCellItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaOfInterestCellItem{}, 0, err
	}
	return v, len(data), err
}
