// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CAGListforMDTItem represents ASN.1 CAGListforMDTItem.
type CAGListforMDTItem struct {
	PlmnID            PLMNIdentity
	CAGID             CAGID
	HasIEExtensions   bool
	IEExtensions      []CAGListforMDTItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CAGListforMDTItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CAGListforMDTItem", v)
}

// DecodeCAGListforMDTItem decodes a standalone APER value and reports the consumed octets.
func DecodeCAGListforMDTItem(data []byte) (CAGListforMDTItem, int, error) {
	var v CAGListforMDTItem
	err := ngapcodec.DecodeParameter("CAGListforMDTItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return CAGListforMDTItem{}, 0, err
	}
	return v, len(data), err
}
