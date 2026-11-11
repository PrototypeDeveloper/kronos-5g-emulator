// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AssistanceDataForRecommendedCellsExtIEsEntryPresent uint32

const (
	AssistanceDataForRecommendedCellsExtIEsEntryPresentNothing AssistanceDataForRecommendedCellsExtIEsEntryPresent = iota
	AssistanceDataForRecommendedCellsExtIEsEntryPresentUnknown
)

type AssistanceDataForRecommendedCellsExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AssistanceDataForRecommendedCellsExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AssistanceDataForRecommendedCellsExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AssistanceDataForRecommendedCells-ExtIEsEntry", v)
}

// DecodeAssistanceDataForRecommendedCellsExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAssistanceDataForRecommendedCellsExtIEsEntry(data []byte) (AssistanceDataForRecommendedCellsExtIEsEntry, int, error) {
	var v AssistanceDataForRecommendedCellsExtIEsEntry
	err := ngapcodec.DecodeParameter("AssistanceDataForRecommendedCells-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AssistanceDataForRecommendedCellsExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
