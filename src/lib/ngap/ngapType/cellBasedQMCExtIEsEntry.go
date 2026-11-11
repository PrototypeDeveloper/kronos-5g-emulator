// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CellBasedQMCExtIEsEntryPresent uint32

const (
	CellBasedQMCExtIEsEntryPresentNothing CellBasedQMCExtIEsEntryPresent = iota
	CellBasedQMCExtIEsEntryPresentUnknown
)

type CellBasedQMCExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CellBasedQMCExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellBasedQMCExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellBasedQMC-ExtIEsEntry", v)
}

// DecodeCellBasedQMCExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCellBasedQMCExtIEsEntry(data []byte) (CellBasedQMCExtIEsEntry, int, error) {
	var v CellBasedQMCExtIEsEntry
	err := ngapcodec.DecodeParameter("CellBasedQMC-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellBasedQMCExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
