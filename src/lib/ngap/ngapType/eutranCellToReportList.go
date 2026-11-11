// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EUTRANCellToReportList represents ASN.1 EUTRAN-CellToReportList.
type EUTRANCellToReportList struct {
	List []EUTRANCellToReportItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRANCellToReportList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRAN-CellToReportList", v)
}

// DecodeEUTRANCellToReportList decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRANCellToReportList(data []byte) (EUTRANCellToReportList, int, error) {
	var v EUTRANCellToReportList
	err := ngapcodec.DecodeParameter("EUTRAN-CellToReportList", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRANCellToReportList{}, 0, err
	}
	return v, len(data), err
}
