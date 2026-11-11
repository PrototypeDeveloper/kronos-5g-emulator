// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RecommendedCellsForPagingExtIEsEntryPresent uint32

const (
	RecommendedCellsForPagingExtIEsEntryPresentNothing RecommendedCellsForPagingExtIEsEntryPresent = iota
	RecommendedCellsForPagingExtIEsEntryPresentUnknown
)

type RecommendedCellsForPagingExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      RecommendedCellsForPagingExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RecommendedCellsForPagingExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RecommendedCellsForPaging-ExtIEsEntry", v)
}

// DecodeRecommendedCellsForPagingExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRecommendedCellsForPagingExtIEsEntry(data []byte) (RecommendedCellsForPagingExtIEsEntry, int, error) {
	var v RecommendedCellsForPagingExtIEsEntry
	err := ngapcodec.DecodeParameter("RecommendedCellsForPaging-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RecommendedCellsForPagingExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
