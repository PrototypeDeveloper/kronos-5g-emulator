// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// OverloadStartNSSAIList represents ASN.1 OverloadStartNSSAIList.
type OverloadStartNSSAIList struct {
	List []OverloadStartNSSAIItem `aper:"sizeLB:1,sizeUB:1024"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v OverloadStartNSSAIList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("OverloadStartNSSAIList", v)
}

// DecodeOverloadStartNSSAIList decodes a standalone APER value and reports the consumed octets.
func DecodeOverloadStartNSSAIList(data []byte) (OverloadStartNSSAIList, int, error) {
	var v OverloadStartNSSAIList
	err := ngapcodec.DecodeParameter("OverloadStartNSSAIList", data, &v)
	if ngapcodec.IsFatal(err) {
		return OverloadStartNSSAIList{}, 0, err
	}
	return v, len(data), err
}
