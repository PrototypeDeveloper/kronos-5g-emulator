// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RecommendedCellList represents ASN.1 RecommendedCellList.
type RecommendedCellList struct {
	List []RecommendedCellItem `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RecommendedCellList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RecommendedCellList", v)
}

// DecodeRecommendedCellList decodes a standalone APER value and reports the consumed octets.
func DecodeRecommendedCellList(data []byte) (RecommendedCellList, int, error) {
	var v RecommendedCellList
	err := ngapcodec.DecodeParameter("RecommendedCellList", data, &v)
	if ngapcodec.IsFatal(err) {
		return RecommendedCellList{}, 0, err
	}
	return v, len(data), err
}
