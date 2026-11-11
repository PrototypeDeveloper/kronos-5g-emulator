// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGRANCellToReportList represents ASN.1 NGRAN-CellToReportList.
type NGRANCellToReportList struct {
	List []NGRANCellToReportItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANCellToReportList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGRAN-CellToReportList", v)
}

// DecodeNGRANCellToReportList decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANCellToReportList(data []byte) (NGRANCellToReportList, int, error) {
	var v NGRANCellToReportList
	err := ngapcodec.DecodeParameter("NGRAN-CellToReportList", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANCellToReportList{}, 0, err
	}
	return v, len(data), err
}
