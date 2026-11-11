// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TargetNSSAIItem represents ASN.1 TargetNSSAI-Item.
type TargetNSSAIItem struct {
	SNSSAI            SNSSAI
	HasIEExtensions   bool
	IEExtensions      []TargetNSSAIItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetNSSAIItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetNSSAI-Item", v)
}

// DecodeTargetNSSAIItem decodes a standalone APER value and reports the consumed octets.
func DecodeTargetNSSAIItem(data []byte) (TargetNSSAIItem, int, error) {
	var v TargetNSSAIItem
	err := ngapcodec.DecodeParameter("TargetNSSAI-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetNSSAIItem{}, 0, err
	}
	return v, len(data), err
}
