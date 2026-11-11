// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RecommendedCellItem represents ASN.1 RecommendedCellItem.
type RecommendedCellItem struct {
	NGRANCGI            NGRANCGI
	HasTimeStayedInCell bool
	TimeStayedInCell    int64 `aper:"valueLB:0,valueUB:4095"`
	HasIEExtensions     bool
	IEExtensions        []RecommendedCellItemExtIEsEntry
	UnknownExtensions   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RecommendedCellItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RecommendedCellItem", v)
}

// DecodeRecommendedCellItem decodes a standalone APER value and reports the consumed octets.
func DecodeRecommendedCellItem(data []byte) (RecommendedCellItem, int, error) {
	var v RecommendedCellItem
	err := ngapcodec.DecodeParameter("RecommendedCellItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return RecommendedCellItem{}, 0, err
	}
	return v, len(data), err
}
