// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GUAMIListItem represents ASN.1 GUAMIList-Item.
type GUAMIListItem struct {
	GUAMI             GUAMI
	HasIEExtensions   bool
	IEExtensions      []GUAMIListItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GUAMIListItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GUAMIList-Item", v)
}

// DecodeGUAMIListItem decodes a standalone APER value and reports the consumed octets.
func DecodeGUAMIListItem(data []byte) (GUAMIListItem, int, error) {
	var v GUAMIListItem
	err := ngapcodec.DecodeParameter("GUAMIList-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return GUAMIListItem{}, 0, err
	}
	return v, len(data), err
}
