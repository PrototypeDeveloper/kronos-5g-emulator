// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAIListForPagingItem represents ASN.1 TAIListForPagingItem.
type TAIListForPagingItem struct {
	TAI               TAI
	HasIEExtensions   bool
	IEExtensions      []TAIListForPagingItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIListForPagingItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAIListForPagingItem", v)
}

// DecodeTAIListForPagingItem decodes a standalone APER value and reports the consumed octets.
func DecodeTAIListForPagingItem(data []byte) (TAIListForPagingItem, int, error) {
	var v TAIListForPagingItem
	err := ngapcodec.DecodeParameter("TAIListForPagingItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIListForPagingItem{}, 0, err
	}
	return v, len(data), err
}
