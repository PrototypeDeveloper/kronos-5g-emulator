// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RecommendedRANNodesForPaging represents ASN.1 RecommendedRANNodesForPaging.
type RecommendedRANNodesForPaging struct {
	RecommendedRANNodeList RecommendedRANNodeList
	HasIEExtensions        bool
	IEExtensions           []RecommendedRANNodesForPagingExtIEsEntry
	UnknownExtensions      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RecommendedRANNodesForPaging) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RecommendedRANNodesForPaging", v)
}

// DecodeRecommendedRANNodesForPaging decodes a standalone APER value and reports the consumed octets.
func DecodeRecommendedRANNodesForPaging(data []byte) (RecommendedRANNodesForPaging, int, error) {
	var v RecommendedRANNodesForPaging
	err := ngapcodec.DecodeParameter("RecommendedRANNodesForPaging", data, &v)
	if ngapcodec.IsFatal(err) {
		return RecommendedRANNodesForPaging{}, 0, err
	}
	return v, len(data), err
}
