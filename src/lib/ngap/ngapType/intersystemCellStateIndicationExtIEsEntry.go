// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type IntersystemCellStateIndicationExtIEsEntryPresent uint32

const (
	IntersystemCellStateIndicationExtIEsEntryPresentNothing IntersystemCellStateIndicationExtIEsEntryPresent = iota
	IntersystemCellStateIndicationExtIEsEntryPresentUnknown
)

type IntersystemCellStateIndicationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      IntersystemCellStateIndicationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemCellStateIndicationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemCellStateIndication-ExtIEsEntry", v)
}

// DecodeIntersystemCellStateIndicationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemCellStateIndicationExtIEsEntry(data []byte) (IntersystemCellStateIndicationExtIEsEntry, int, error) {
	var v IntersystemCellStateIndicationExtIEsEntry
	err := ngapcodec.DecodeParameter("IntersystemCellStateIndication-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemCellStateIndicationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
