// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EUTRANCellToReportItem represents ASN.1 EUTRAN-CellToReportItem.
type EUTRANCellToReportItem struct {
	ECGI              EUTRACGI
	HasIEExtensions   bool
	IEExtensions      []EUTRANCellToReportItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRANCellToReportItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRAN-CellToReportItem", v)
}

// DecodeEUTRANCellToReportItem decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRANCellToReportItem(data []byte) (EUTRANCellToReportItem, int, error) {
	var v EUTRANCellToReportItem
	err := ngapcodec.DecodeParameter("EUTRAN-CellToReportItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRANCellToReportItem{}, 0, err
	}
	return v, len(data), err
}
