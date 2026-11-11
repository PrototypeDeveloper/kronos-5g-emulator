// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AreaOfInterestTAIItem represents ASN.1 AreaOfInterestTAIItem.
type AreaOfInterestTAIItem struct {
	TAI               TAI
	HasIEExtensions   bool
	IEExtensions      []AreaOfInterestTAIItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaOfInterestTAIItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaOfInterestTAIItem", v)
}

// DecodeAreaOfInterestTAIItem decodes a standalone APER value and reports the consumed octets.
func DecodeAreaOfInterestTAIItem(data []byte) (AreaOfInterestTAIItem, int, error) {
	var v AreaOfInterestTAIItem
	err := ngapcodec.DecodeParameter("AreaOfInterestTAIItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaOfInterestTAIItem{}, 0, err
	}
	return v, len(data), err
}
