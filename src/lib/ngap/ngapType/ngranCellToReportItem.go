// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGRANCellToReportItem represents ASN.1 NGRAN-CellToReportItem.
type NGRANCellToReportItem struct {
	NGRANCGI          NGRANCGI
	HasIEExtensions   bool
	IEExtensions      []NGRANCellToReportItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANCellToReportItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGRAN-CellToReportItem", v)
}

// DecodeNGRANCellToReportItem decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANCellToReportItem(data []byte) (NGRANCellToReportItem, int, error) {
	var v NGRANCellToReportItem
	err := ngapcodec.DecodeParameter("NGRAN-CellToReportItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANCellToReportItem{}, 0, err
	}
	return v, len(data), err
}
