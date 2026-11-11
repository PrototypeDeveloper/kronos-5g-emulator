// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type InfoOnRecommendedCellsAndRANNodesForPagingExtIEsEntryPresent uint32

const (
	InfoOnRecommendedCellsAndRANNodesForPagingExtIEsEntryPresentNothing InfoOnRecommendedCellsAndRANNodesForPagingExtIEsEntryPresent = iota
	InfoOnRecommendedCellsAndRANNodesForPagingExtIEsEntryPresentUnknown
)

type InfoOnRecommendedCellsAndRANNodesForPagingExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      InfoOnRecommendedCellsAndRANNodesForPagingExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InfoOnRecommendedCellsAndRANNodesForPagingExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InfoOnRecommendedCellsAndRANNodesForPaging-ExtIEsEntry", v)
}

// DecodeInfoOnRecommendedCellsAndRANNodesForPagingExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeInfoOnRecommendedCellsAndRANNodesForPagingExtIEsEntry(data []byte) (InfoOnRecommendedCellsAndRANNodesForPagingExtIEsEntry, int, error) {
	var v InfoOnRecommendedCellsAndRANNodesForPagingExtIEsEntry
	err := ngapcodec.DecodeParameter("InfoOnRecommendedCellsAndRANNodesForPaging-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return InfoOnRecommendedCellsAndRANNodesForPagingExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
