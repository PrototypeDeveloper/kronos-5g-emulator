// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CompletedCellsInEAINRItemExtIEsEntryPresent uint32

const (
	CompletedCellsInEAINRItemExtIEsEntryPresentNothing CompletedCellsInEAINRItemExtIEsEntryPresent = iota
	CompletedCellsInEAINRItemExtIEsEntryPresentUnknown
)

type CompletedCellsInEAINRItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CompletedCellsInEAINRItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CompletedCellsInEAINRItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CompletedCellsInEAI-NR-Item-ExtIEsEntry", v)
}

// DecodeCompletedCellsInEAINRItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCompletedCellsInEAINRItemExtIEsEntry(data []byte) (CompletedCellsInEAINRItemExtIEsEntry, int, error) {
	var v CompletedCellsInEAINRItemExtIEsEntry
	err := ngapcodec.DecodeParameter("CompletedCellsInEAI-NR-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CompletedCellsInEAINRItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
