// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CellBasedMDTNRExtIEsEntryPresent uint32

const (
	CellBasedMDTNRExtIEsEntryPresentNothing CellBasedMDTNRExtIEsEntryPresent = iota
	CellBasedMDTNRExtIEsEntryPresentUnknown
)

type CellBasedMDTNRExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CellBasedMDTNRExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellBasedMDTNRExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellBasedMDT-NR-ExtIEsEntry", v)
}

// DecodeCellBasedMDTNRExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCellBasedMDTNRExtIEsEntry(data []byte) (CellBasedMDTNRExtIEsEntry, int, error) {
	var v CellBasedMDTNRExtIEsEntry
	err := ngapcodec.DecodeParameter("CellBasedMDT-NR-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellBasedMDTNRExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
