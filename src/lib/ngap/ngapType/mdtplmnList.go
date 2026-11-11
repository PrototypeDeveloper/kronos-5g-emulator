// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MDTPLMNList represents ASN.1 MDTPLMNList.
type MDTPLMNList struct {
	List []PLMNIdentity `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTPLMNList) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("MDTPLMNList", v) }

// DecodeMDTPLMNList decodes a standalone APER value and reports the consumed octets.
func DecodeMDTPLMNList(data []byte) (MDTPLMNList, int, error) {
	var v MDTPLMNList
	err := ngapcodec.DecodeParameter("MDTPLMNList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTPLMNList{}, 0, err
	}
	return v, len(data), err
}
