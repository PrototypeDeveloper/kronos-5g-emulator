// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RecommendedRANNodeItemExtIEsEntryPresent uint32

const (
	RecommendedRANNodeItemExtIEsEntryPresentNothing RecommendedRANNodeItemExtIEsEntryPresent = iota
	RecommendedRANNodeItemExtIEsEntryPresentUnknown
)

type RecommendedRANNodeItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      RecommendedRANNodeItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RecommendedRANNodeItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RecommendedRANNodeItem-ExtIEsEntry", v)
}

// DecodeRecommendedRANNodeItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRecommendedRANNodeItemExtIEsEntry(data []byte) (RecommendedRANNodeItemExtIEsEntry, int, error) {
	var v RecommendedRANNodeItemExtIEsEntry
	err := ngapcodec.DecodeParameter("RecommendedRANNodeItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RecommendedRANNodeItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
