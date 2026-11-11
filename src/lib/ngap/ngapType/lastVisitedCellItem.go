// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LastVisitedCellItem represents ASN.1 LastVisitedCellItem.
type LastVisitedCellItem struct {
	LastVisitedCellInformation LastVisitedCellInformation
	HasIEExtensions            bool
	IEExtensions               []LastVisitedCellItemExtIEsEntry
	UnknownExtensions          []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LastVisitedCellItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LastVisitedCellItem", v)
}

// DecodeLastVisitedCellItem decodes a standalone APER value and reports the consumed octets.
func DecodeLastVisitedCellItem(data []byte) (LastVisitedCellItem, int, error) {
	var v LastVisitedCellItem
	err := ngapcodec.DecodeParameter("LastVisitedCellItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return LastVisitedCellItem{}, 0, err
	}
	return v, len(data), err
}
