// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CompletedCellsInTAIEUTRAItemExtIEsEntryPresent uint32

const (
	CompletedCellsInTAIEUTRAItemExtIEsEntryPresentNothing CompletedCellsInTAIEUTRAItemExtIEsEntryPresent = iota
	CompletedCellsInTAIEUTRAItemExtIEsEntryPresentUnknown
)

type CompletedCellsInTAIEUTRAItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CompletedCellsInTAIEUTRAItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CompletedCellsInTAIEUTRAItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CompletedCellsInTAI-EUTRA-Item-ExtIEsEntry", v)
}

// DecodeCompletedCellsInTAIEUTRAItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCompletedCellsInTAIEUTRAItemExtIEsEntry(data []byte) (CompletedCellsInTAIEUTRAItemExtIEsEntry, int, error) {
	var v CompletedCellsInTAIEUTRAItemExtIEsEntry
	err := ngapcodec.DecodeParameter("CompletedCellsInTAI-EUTRA-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CompletedCellsInTAIEUTRAItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
