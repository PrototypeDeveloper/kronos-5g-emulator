// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SupportedTAItem represents ASN.1 SupportedTAItem.
type SupportedTAItem struct {
	TAC               TAC
	BroadcastPLMNList BroadcastPLMNList
	HasIEExtensions   bool
	IEExtensions      []SupportedTAItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SupportedTAItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SupportedTAItem", v)
}

// DecodeSupportedTAItem decodes a standalone APER value and reports the consumed octets.
func DecodeSupportedTAItem(data []byte) (SupportedTAItem, int, error) {
	var v SupportedTAItem
	err := ngapcodec.DecodeParameter("SupportedTAItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return SupportedTAItem{}, 0, err
	}
	return v, len(data), err
}
