// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AllowedNSSAIItem represents ASN.1 AllowedNSSAI-Item.
type AllowedNSSAIItem struct {
	SNSSAI            SNSSAI
	HasIEExtensions   bool
	IEExtensions      []AllowedNSSAIItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AllowedNSSAIItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AllowedNSSAI-Item", v)
}

// DecodeAllowedNSSAIItem decodes a standalone APER value and reports the consumed octets.
func DecodeAllowedNSSAIItem(data []byte) (AllowedNSSAIItem, int, error) {
	var v AllowedNSSAIItem
	err := ngapcodec.DecodeParameter("AllowedNSSAI-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return AllowedNSSAIItem{}, 0, err
	}
	return v, len(data), err
}
