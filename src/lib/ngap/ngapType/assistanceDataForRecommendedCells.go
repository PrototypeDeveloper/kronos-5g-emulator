// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AssistanceDataForRecommendedCells represents ASN.1 AssistanceDataForRecommendedCells.
type AssistanceDataForRecommendedCells struct {
	RecommendedCellsForPaging RecommendedCellsForPaging
	HasIEExtensions           bool
	IEExtensions              []AssistanceDataForRecommendedCellsExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AssistanceDataForRecommendedCells) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AssistanceDataForRecommendedCells", v)
}

// DecodeAssistanceDataForRecommendedCells decodes a standalone APER value and reports the consumed octets.
func DecodeAssistanceDataForRecommendedCells(data []byte) (AssistanceDataForRecommendedCells, int, error) {
	var v AssistanceDataForRecommendedCells
	err := ngapcodec.DecodeParameter("AssistanceDataForRecommendedCells", data, &v)
	if ngapcodec.IsFatal(err) {
		return AssistanceDataForRecommendedCells{}, 0, err
	}
	return v, len(data), err
}
