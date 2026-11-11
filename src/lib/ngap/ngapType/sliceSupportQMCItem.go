// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SliceSupportQMCItem represents ASN.1 SliceSupportQMC-Item.
type SliceSupportQMCItem struct {
	SNSSAI            SNSSAI
	HasIEExtensions   bool
	IEExtensions      []SliceSupportQMCItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SliceSupportQMCItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SliceSupportQMC-Item", v)
}

// DecodeSliceSupportQMCItem decodes a standalone APER value and reports the consumed octets.
func DecodeSliceSupportQMCItem(data []byte) (SliceSupportQMCItem, int, error) {
	var v SliceSupportQMCItem
	err := ngapcodec.DecodeParameter("SliceSupportQMC-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return SliceSupportQMCItem{}, 0, err
	}
	return v, len(data), err
}
