// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RecommendedCellsForPaging represents ASN.1 RecommendedCellsForPaging.
type RecommendedCellsForPaging struct {
	RecommendedCellList RecommendedCellList
	HasIEExtensions     bool
	IEExtensions        []RecommendedCellsForPagingExtIEsEntry
	UnknownExtensions   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RecommendedCellsForPaging) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RecommendedCellsForPaging", v)
}

// DecodeRecommendedCellsForPaging decodes a standalone APER value and reports the consumed octets.
func DecodeRecommendedCellsForPaging(data []byte) (RecommendedCellsForPaging, int, error) {
	var v RecommendedCellsForPaging
	err := ngapcodec.DecodeParameter("RecommendedCellsForPaging", data, &v)
	if ngapcodec.IsFatal(err) {
		return RecommendedCellsForPaging{}, 0, err
	}
	return v, len(data), err
}
