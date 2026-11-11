// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MDTPLMNModificationList represents ASN.1 MDTPLMNModificationList.
type MDTPLMNModificationList struct {
	List []PLMNIdentity `aper:"sizeLB:0,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTPLMNModificationList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MDTPLMNModificationList", v)
}

// DecodeMDTPLMNModificationList decodes a standalone APER value and reports the consumed octets.
func DecodeMDTPLMNModificationList(data []byte) (MDTPLMNModificationList, int, error) {
	var v MDTPLMNModificationList
	err := ngapcodec.DecodeParameter("MDTPLMNModificationList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTPLMNModificationList{}, 0, err
	}
	return v, len(data), err
}
