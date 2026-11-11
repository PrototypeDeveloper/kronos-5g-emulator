// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CancelledCellsInEAIEUTRAItemExtIEsEntryPresent uint32

const (
	CancelledCellsInEAIEUTRAItemExtIEsEntryPresentNothing CancelledCellsInEAIEUTRAItemExtIEsEntryPresent = iota
	CancelledCellsInEAIEUTRAItemExtIEsEntryPresentUnknown
)

type CancelledCellsInEAIEUTRAItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CancelledCellsInEAIEUTRAItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CancelledCellsInEAIEUTRAItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CancelledCellsInEAI-EUTRA-Item-ExtIEsEntry", v)
}

// DecodeCancelledCellsInEAIEUTRAItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCancelledCellsInEAIEUTRAItemExtIEsEntry(data []byte) (CancelledCellsInEAIEUTRAItemExtIEsEntry, int, error) {
	var v CancelledCellsInEAIEUTRAItemExtIEsEntry
	err := ngapcodec.DecodeParameter("CancelledCellsInEAI-EUTRA-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CancelledCellsInEAIEUTRAItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
