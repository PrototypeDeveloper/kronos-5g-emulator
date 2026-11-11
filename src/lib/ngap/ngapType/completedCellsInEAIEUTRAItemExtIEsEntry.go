// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CompletedCellsInEAIEUTRAItemExtIEsEntryPresent uint32

const (
	CompletedCellsInEAIEUTRAItemExtIEsEntryPresentNothing CompletedCellsInEAIEUTRAItemExtIEsEntryPresent = iota
	CompletedCellsInEAIEUTRAItemExtIEsEntryPresentUnknown
)

type CompletedCellsInEAIEUTRAItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CompletedCellsInEAIEUTRAItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CompletedCellsInEAIEUTRAItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CompletedCellsInEAI-EUTRA-Item-ExtIEsEntry", v)
}

// DecodeCompletedCellsInEAIEUTRAItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCompletedCellsInEAIEUTRAItemExtIEsEntry(data []byte) (CompletedCellsInEAIEUTRAItemExtIEsEntry, int, error) {
	var v CompletedCellsInEAIEUTRAItemExtIEsEntry
	err := ngapcodec.DecodeParameter("CompletedCellsInEAI-EUTRA-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CompletedCellsInEAIEUTRAItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
