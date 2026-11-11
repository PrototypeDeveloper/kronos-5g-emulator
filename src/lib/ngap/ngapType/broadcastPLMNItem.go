// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BroadcastPLMNItem represents ASN.1 BroadcastPLMNItem.
type BroadcastPLMNItem struct {
	PLMNIdentity        PLMNIdentity
	TAISliceSupportList SliceSupportList
	HasIEExtensions     bool
	IEExtensions        []BroadcastPLMNItemExtIEsEntry
	UnknownExtensions   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BroadcastPLMNItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BroadcastPLMNItem", v)
}

// DecodeBroadcastPLMNItem decodes a standalone APER value and reports the consumed octets.
func DecodeBroadcastPLMNItem(data []byte) (BroadcastPLMNItem, int, error) {
	var v BroadcastPLMNItem
	err := ngapcodec.DecodeParameter("BroadcastPLMNItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return BroadcastPLMNItem{}, 0, err
	}
	return v, len(data), err
}
