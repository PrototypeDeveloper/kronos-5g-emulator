// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RequestedNSSAIItem represents ASN.1 RequestedNSSAIItem.
type RequestedNSSAIItem struct {
	SNSSAI            SNSSAI
	HasIEExtensions   bool
	IEExtensions      []RequestedNSSAIItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RequestedNSSAIItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RequestedNSSAIItem", v)
}

// DecodeRequestedNSSAIItem decodes a standalone APER value and reports the consumed octets.
func DecodeRequestedNSSAIItem(data []byte) (RequestedNSSAIItem, int, error) {
	var v RequestedNSSAIItem
	err := ngapcodec.DecodeParameter("RequestedNSSAIItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return RequestedNSSAIItem{}, 0, err
	}
	return v, len(data), err
}
