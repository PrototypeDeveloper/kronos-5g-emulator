// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RecommendedRANNodesForPagingExtIEsEntryPresent uint32

const (
	RecommendedRANNodesForPagingExtIEsEntryPresentNothing RecommendedRANNodesForPagingExtIEsEntryPresent = iota
	RecommendedRANNodesForPagingExtIEsEntryPresentUnknown
)

type RecommendedRANNodesForPagingExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      RecommendedRANNodesForPagingExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RecommendedRANNodesForPagingExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RecommendedRANNodesForPaging-ExtIEsEntry", v)
}

// DecodeRecommendedRANNodesForPagingExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRecommendedRANNodesForPagingExtIEsEntry(data []byte) (RecommendedRANNodesForPagingExtIEsEntry, int, error) {
	var v RecommendedRANNodesForPagingExtIEsEntry
	err := ngapcodec.DecodeParameter("RecommendedRANNodesForPaging-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RecommendedRANNodesForPagingExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
