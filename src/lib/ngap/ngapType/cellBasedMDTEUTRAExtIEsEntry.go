// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CellBasedMDTEUTRAExtIEsEntryPresent uint32

const (
	CellBasedMDTEUTRAExtIEsEntryPresentNothing CellBasedMDTEUTRAExtIEsEntryPresent = iota
	CellBasedMDTEUTRAExtIEsEntryPresentUnknown
)

type CellBasedMDTEUTRAExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CellBasedMDTEUTRAExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellBasedMDTEUTRAExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellBasedMDT-EUTRA-ExtIEsEntry", v)
}

// DecodeCellBasedMDTEUTRAExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCellBasedMDTEUTRAExtIEsEntry(data []byte) (CellBasedMDTEUTRAExtIEsEntry, int, error) {
	var v CellBasedMDTEUTRAExtIEsEntry
	err := ngapcodec.DecodeParameter("CellBasedMDT-EUTRA-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellBasedMDTEUTRAExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
