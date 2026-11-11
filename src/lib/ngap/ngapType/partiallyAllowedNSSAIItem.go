// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PartiallyAllowedNSSAIItem represents ASN.1 PartiallyAllowedNSSAI-Item.
type PartiallyAllowedNSSAIItem struct {
	SNSSAI            SNSSAI
	HasIEExtensions   bool
	IEExtensions      []PartiallyAllowedNSSAIItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PartiallyAllowedNSSAIItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PartiallyAllowedNSSAI-Item", v)
}

// DecodePartiallyAllowedNSSAIItem decodes a standalone APER value and reports the consumed octets.
func DecodePartiallyAllowedNSSAIItem(data []byte) (PartiallyAllowedNSSAIItem, int, error) {
	var v PartiallyAllowedNSSAIItem
	err := ngapcodec.DecodeParameter("PartiallyAllowedNSSAI-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return PartiallyAllowedNSSAIItem{}, 0, err
	}
	return v, len(data), err
}
