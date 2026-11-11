// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CancelledCellsInTAINRItemExtIEsEntryPresent uint32

const (
	CancelledCellsInTAINRItemExtIEsEntryPresentNothing CancelledCellsInTAINRItemExtIEsEntryPresent = iota
	CancelledCellsInTAINRItemExtIEsEntryPresentUnknown
)

type CancelledCellsInTAINRItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CancelledCellsInTAINRItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CancelledCellsInTAINRItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CancelledCellsInTAI-NR-Item-ExtIEsEntry", v)
}

// DecodeCancelledCellsInTAINRItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCancelledCellsInTAINRItemExtIEsEntry(data []byte) (CancelledCellsInTAINRItemExtIEsEntry, int, error) {
	var v CancelledCellsInTAINRItemExtIEsEntry
	err := ngapcodec.DecodeParameter("CancelledCellsInTAI-NR-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CancelledCellsInTAINRItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
