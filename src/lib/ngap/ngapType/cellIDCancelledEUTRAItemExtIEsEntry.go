// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CellIDCancelledEUTRAItemExtIEsEntryPresent uint32

const (
	CellIDCancelledEUTRAItemExtIEsEntryPresentNothing CellIDCancelledEUTRAItemExtIEsEntryPresent = iota
	CellIDCancelledEUTRAItemExtIEsEntryPresentUnknown
)

type CellIDCancelledEUTRAItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CellIDCancelledEUTRAItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellIDCancelledEUTRAItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellIDCancelledEUTRA-Item-ExtIEsEntry", v)
}

// DecodeCellIDCancelledEUTRAItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCellIDCancelledEUTRAItemExtIEsEntry(data []byte) (CellIDCancelledEUTRAItemExtIEsEntry, int, error) {
	var v CellIDCancelledEUTRAItemExtIEsEntry
	err := ngapcodec.DecodeParameter("CellIDCancelledEUTRA-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellIDCancelledEUTRAItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
