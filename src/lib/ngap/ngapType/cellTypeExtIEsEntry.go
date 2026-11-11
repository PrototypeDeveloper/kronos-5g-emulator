// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CellTypeExtIEsEntryPresent uint32

const (
	CellTypeExtIEsEntryPresentNothing CellTypeExtIEsEntryPresent = iota
	CellTypeExtIEsEntryPresentUnknown
)

type CellTypeExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CellTypeExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellTypeExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellType-ExtIEsEntry", v)
}

// DecodeCellTypeExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCellTypeExtIEsEntry(data []byte) (CellTypeExtIEsEntry, int, error) {
	var v CellTypeExtIEsEntry
	err := ngapcodec.DecodeParameter("CellType-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellTypeExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
