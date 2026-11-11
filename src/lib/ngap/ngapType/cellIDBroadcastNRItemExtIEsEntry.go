// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CellIDBroadcastNRItemExtIEsEntryPresent uint32

const (
	CellIDBroadcastNRItemExtIEsEntryPresentNothing CellIDBroadcastNRItemExtIEsEntryPresent = iota
	CellIDBroadcastNRItemExtIEsEntryPresentUnknown
)

type CellIDBroadcastNRItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CellIDBroadcastNRItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellIDBroadcastNRItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellIDBroadcastNR-Item-ExtIEsEntry", v)
}

// DecodeCellIDBroadcastNRItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCellIDBroadcastNRItemExtIEsEntry(data []byte) (CellIDBroadcastNRItemExtIEsEntry, int, error) {
	var v CellIDBroadcastNRItemExtIEsEntry
	err := ngapcodec.DecodeParameter("CellIDBroadcastNR-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellIDBroadcastNRItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
