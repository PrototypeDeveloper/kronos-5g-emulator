// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAIMBSSupportItem represents ASN.1 TAIMBSSupportItem.
type TAIMBSSupportItem struct {
	TAI               TAI
	HasIEExtensions   bool
	IEExtensions      []TAIMBSSupportItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIMBSSupportItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAIMBSSupportItem", v)
}

// DecodeTAIMBSSupportItem decodes a standalone APER value and reports the consumed octets.
func DecodeTAIMBSSupportItem(data []byte) (TAIMBSSupportItem, int, error) {
	var v TAIMBSSupportItem
	err := ngapcodec.DecodeParameter("TAIMBSSupportItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIMBSSupportItem{}, 0, err
	}
	return v, len(data), err
}
