// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAIListForInactiveItem represents ASN.1 TAIListForInactiveItem.
type TAIListForInactiveItem struct {
	TAI               TAI
	HasIEExtensions   bool
	IEExtensions      []TAIListForInactiveItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIListForInactiveItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAIListForInactiveItem", v)
}

// DecodeTAIListForInactiveItem decodes a standalone APER value and reports the consumed octets.
func DecodeTAIListForInactiveItem(data []byte) (TAIListForInactiveItem, int, error) {
	var v TAIListForInactiveItem
	err := ngapcodec.DecodeParameter("TAIListForInactiveItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIListForInactiveItem{}, 0, err
	}
	return v, len(data), err
}
