// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SliceOverloadItem represents ASN.1 SliceOverloadItem.
type SliceOverloadItem struct {
	SNSSAI            SNSSAI
	HasIEExtensions   bool
	IEExtensions      []SliceOverloadItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SliceOverloadItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SliceOverloadItem", v)
}

// DecodeSliceOverloadItem decodes a standalone APER value and reports the consumed octets.
func DecodeSliceOverloadItem(data []byte) (SliceOverloadItem, int, error) {
	var v SliceOverloadItem
	err := ngapcodec.DecodeParameter("SliceOverloadItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return SliceOverloadItem{}, 0, err
	}
	return v, len(data), err
}
