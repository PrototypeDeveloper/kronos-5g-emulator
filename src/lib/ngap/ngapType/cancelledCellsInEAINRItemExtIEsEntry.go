// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CancelledCellsInEAINRItemExtIEsEntryPresent uint32

const (
	CancelledCellsInEAINRItemExtIEsEntryPresentNothing CancelledCellsInEAINRItemExtIEsEntryPresent = iota
	CancelledCellsInEAINRItemExtIEsEntryPresentUnknown
)

type CancelledCellsInEAINRItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CancelledCellsInEAINRItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CancelledCellsInEAINRItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CancelledCellsInEAI-NR-Item-ExtIEsEntry", v)
}

// DecodeCancelledCellsInEAINRItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCancelledCellsInEAINRItemExtIEsEntry(data []byte) (CancelledCellsInEAINRItemExtIEsEntry, int, error) {
	var v CancelledCellsInEAINRItemExtIEsEntry
	err := ngapcodec.DecodeParameter("CancelledCellsInEAI-NR-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CancelledCellsInEAINRItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
