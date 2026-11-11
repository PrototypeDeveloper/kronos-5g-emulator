// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CompletedCellsInTAINRItemExtIEsEntryPresent uint32

const (
	CompletedCellsInTAINRItemExtIEsEntryPresentNothing CompletedCellsInTAINRItemExtIEsEntryPresent = iota
	CompletedCellsInTAINRItemExtIEsEntryPresentUnknown
)

type CompletedCellsInTAINRItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CompletedCellsInTAINRItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CompletedCellsInTAINRItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CompletedCellsInTAI-NR-Item-ExtIEsEntry", v)
}

// DecodeCompletedCellsInTAINRItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCompletedCellsInTAINRItemExtIEsEntry(data []byte) (CompletedCellsInTAINRItemExtIEsEntry, int, error) {
	var v CompletedCellsInTAINRItemExtIEsEntry
	err := ngapcodec.DecodeParameter("CompletedCellsInTAI-NR-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CompletedCellsInTAINRItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
