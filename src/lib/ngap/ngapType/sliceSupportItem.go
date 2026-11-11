// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SliceSupportItem represents ASN.1 SliceSupportItem.
type SliceSupportItem struct {
	SNSSAI            SNSSAI
	HasIEExtensions   bool
	IEExtensions      []SliceSupportItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SliceSupportItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SliceSupportItem", v)
}

// DecodeSliceSupportItem decodes a standalone APER value and reports the consumed octets.
func DecodeSliceSupportItem(data []byte) (SliceSupportItem, int, error) {
	var v SliceSupportItem
	err := ngapcodec.DecodeParameter("SliceSupportItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return SliceSupportItem{}, 0, err
	}
	return v, len(data), err
}
