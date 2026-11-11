// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RecommendedRANNodeList represents ASN.1 RecommendedRANNodeList.
type RecommendedRANNodeList struct {
	List []RecommendedRANNodeItem `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RecommendedRANNodeList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RecommendedRANNodeList", v)
}

// DecodeRecommendedRANNodeList decodes a standalone APER value and reports the consumed octets.
func DecodeRecommendedRANNodeList(data []byte) (RecommendedRANNodeList, int, error) {
	var v RecommendedRANNodeList
	err := ngapcodec.DecodeParameter("RecommendedRANNodeList", data, &v)
	if ngapcodec.IsFatal(err) {
		return RecommendedRANNodeList{}, 0, err
	}
	return v, len(data), err
}
