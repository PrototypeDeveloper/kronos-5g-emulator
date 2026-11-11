// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EUTRANCellToReportItemExtIEsEntryPresent uint32

const (
	EUTRANCellToReportItemExtIEsEntryPresentNothing EUTRANCellToReportItemExtIEsEntryPresent = iota
	EUTRANCellToReportItemExtIEsEntryPresentUnknown
)

type EUTRANCellToReportItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EUTRANCellToReportItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRANCellToReportItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRAN-CellToReportItem-ExtIEsEntry", v)
}

// DecodeEUTRANCellToReportItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRANCellToReportItemExtIEsEntry(data []byte) (EUTRANCellToReportItemExtIEsEntry, int, error) {
	var v EUTRANCellToReportItemExtIEsEntry
	err := ngapcodec.DecodeParameter("EUTRAN-CellToReportItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRANCellToReportItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
