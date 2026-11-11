// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RecommendedCellItemExtIEsEntryPresent uint32

const (
	RecommendedCellItemExtIEsEntryPresentNothing RecommendedCellItemExtIEsEntryPresent = iota
	RecommendedCellItemExtIEsEntryPresentUnknown
)

type RecommendedCellItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      RecommendedCellItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RecommendedCellItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RecommendedCellItem-ExtIEsEntry", v)
}

// DecodeRecommendedCellItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRecommendedCellItemExtIEsEntry(data []byte) (RecommendedCellItemExtIEsEntry, int, error) {
	var v RecommendedCellItemExtIEsEntry
	err := ngapcodec.DecodeParameter("RecommendedCellItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RecommendedCellItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
