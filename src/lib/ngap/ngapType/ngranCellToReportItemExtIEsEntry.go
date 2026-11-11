// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NGRANCellToReportItemExtIEsEntryPresent uint32

const (
	NGRANCellToReportItemExtIEsEntryPresentNothing NGRANCellToReportItemExtIEsEntryPresent = iota
	NGRANCellToReportItemExtIEsEntryPresentUnknown
)

type NGRANCellToReportItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NGRANCellToReportItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANCellToReportItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGRAN-CellToReportItem-ExtIEsEntry", v)
}

// DecodeNGRANCellToReportItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANCellToReportItemExtIEsEntry(data []byte) (NGRANCellToReportItemExtIEsEntry, int, error) {
	var v NGRANCellToReportItemExtIEsEntry
	err := ngapcodec.DecodeParameter("NGRAN-CellToReportItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANCellToReportItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
