// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SliceMDTItem represents ASN.1 SliceMDTItem.
type SliceMDTItem struct {
	SNSSAI            SNSSAI
	HasIEExtensions   bool
	IEExtensions      []SliceMDTItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SliceMDTItem) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("SliceMDTItem", v) }

// DecodeSliceMDTItem decodes a standalone APER value and reports the consumed octets.
func DecodeSliceMDTItem(data []byte) (SliceMDTItem, int, error) {
	var v SliceMDTItem
	err := ngapcodec.DecodeParameter("SliceMDTItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return SliceMDTItem{}, 0, err
	}
	return v, len(data), err
}
