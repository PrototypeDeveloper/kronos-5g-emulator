// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// InfoOnRecommendedCellsAndRANNodesForPaging represents ASN.1 InfoOnRecommendedCellsAndRANNodesForPaging.
type InfoOnRecommendedCellsAndRANNodesForPaging struct {
	RecommendedCellsForPaging  RecommendedCellsForPaging
	RecommendRANNodesForPaging RecommendedRANNodesForPaging
	HasIEExtensions            bool
	IEExtensions               []InfoOnRecommendedCellsAndRANNodesForPagingExtIEsEntry
	UnknownExtensions          []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InfoOnRecommendedCellsAndRANNodesForPaging) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InfoOnRecommendedCellsAndRANNodesForPaging", v)
}

// DecodeInfoOnRecommendedCellsAndRANNodesForPaging decodes a standalone APER value and reports the consumed octets.
func DecodeInfoOnRecommendedCellsAndRANNodesForPaging(data []byte) (InfoOnRecommendedCellsAndRANNodesForPaging, int, error) {
	var v InfoOnRecommendedCellsAndRANNodesForPaging
	err := ngapcodec.DecodeParameter("InfoOnRecommendedCellsAndRANNodesForPaging", data, &v)
	if ngapcodec.IsFatal(err) {
		return InfoOnRecommendedCellsAndRANNodesForPaging{}, 0, err
	}
	return v, len(data), err
}
