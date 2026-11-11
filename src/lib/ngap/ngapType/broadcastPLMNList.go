// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BroadcastPLMNList represents ASN.1 BroadcastPLMNList.
type BroadcastPLMNList struct {
	List []BroadcastPLMNItem `aper:"sizeLB:1,sizeUB:12"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BroadcastPLMNList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BroadcastPLMNList", v)
}

// DecodeBroadcastPLMNList decodes a standalone APER value and reports the consumed octets.
func DecodeBroadcastPLMNList(data []byte) (BroadcastPLMNList, int, error) {
	var v BroadcastPLMNList
	err := ngapcodec.DecodeParameter("BroadcastPLMNList", data, &v)
	if ngapcodec.IsFatal(err) {
		return BroadcastPLMNList{}, 0, err
	}
	return v, len(data), err
}
