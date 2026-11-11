// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CellIDBroadcastEUTRAItemExtIEsEntryPresent uint32

const (
	CellIDBroadcastEUTRAItemExtIEsEntryPresentNothing CellIDBroadcastEUTRAItemExtIEsEntryPresent = iota
	CellIDBroadcastEUTRAItemExtIEsEntryPresentUnknown
)

type CellIDBroadcastEUTRAItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CellIDBroadcastEUTRAItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellIDBroadcastEUTRAItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellIDBroadcastEUTRA-Item-ExtIEsEntry", v)
}

// DecodeCellIDBroadcastEUTRAItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCellIDBroadcastEUTRAItemExtIEsEntry(data []byte) (CellIDBroadcastEUTRAItemExtIEsEntry, int, error) {
	var v CellIDBroadcastEUTRAItemExtIEsEntry
	err := ngapcodec.DecodeParameter("CellIDBroadcastEUTRA-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellIDBroadcastEUTRAItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
