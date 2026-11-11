// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RecommendedRANNodeItem represents ASN.1 RecommendedRANNodeItem.
type RecommendedRANNodeItem struct {
	AMFPagingTarget   AMFPagingTarget
	HasIEExtensions   bool
	IEExtensions      []RecommendedRANNodeItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RecommendedRANNodeItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RecommendedRANNodeItem", v)
}

// DecodeRecommendedRANNodeItem decodes a standalone APER value and reports the consumed octets.
func DecodeRecommendedRANNodeItem(data []byte) (RecommendedRANNodeItem, int, error) {
	var v RecommendedRANNodeItem
	err := ngapcodec.DecodeParameter("RecommendedRANNodeItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return RecommendedRANNodeItem{}, 0, err
	}
	return v, len(data), err
}
