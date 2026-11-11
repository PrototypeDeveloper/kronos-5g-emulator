// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CellIDListForRestartExtIEsEntryPresent uint32

const (
	CellIDListForRestartExtIEsEntryPresentNothing CellIDListForRestartExtIEsEntryPresent = iota
	CellIDListForRestartExtIEsEntryPresentUnknown
)

type CellIDListForRestartExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      CellIDListForRestartExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellIDListForRestartExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellIDListForRestart-ExtIEsEntry", v)
}

// DecodeCellIDListForRestartExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCellIDListForRestartExtIEsEntry(data []byte) (CellIDListForRestartExtIEsEntry, int, error) {
	var v CellIDListForRestartExtIEsEntry
	err := ngapcodec.DecodeParameter("CellIDListForRestart-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellIDListForRestartExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
