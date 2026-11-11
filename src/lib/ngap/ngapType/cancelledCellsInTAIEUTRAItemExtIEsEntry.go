// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CancelledCellsInTAIEUTRAItemExtIEsEntryPresent uint32

const (
	CancelledCellsInTAIEUTRAItemExtIEsEntryPresentNothing CancelledCellsInTAIEUTRAItemExtIEsEntryPresent = iota
	CancelledCellsInTAIEUTRAItemExtIEsEntryPresentUnknown
)

type CancelledCellsInTAIEUTRAItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CancelledCellsInTAIEUTRAItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CancelledCellsInTAIEUTRAItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CancelledCellsInTAI-EUTRA-Item-ExtIEsEntry", v)
}

// DecodeCancelledCellsInTAIEUTRAItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCancelledCellsInTAIEUTRAItemExtIEsEntry(data []byte) (CancelledCellsInTAIEUTRAItemExtIEsEntry, int, error) {
	var v CancelledCellsInTAIEUTRAItemExtIEsEntry
	err := ngapcodec.DecodeParameter("CancelledCellsInTAI-EUTRA-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CancelledCellsInTAIEUTRAItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
