// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGRANCellReportList represents ASN.1 NGRAN-CellReportList.
type NGRANCellReportList struct {
	List []NGRANCellReportItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANCellReportList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGRAN-CellReportList", v)
}

// DecodeNGRANCellReportList decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANCellReportList(data []byte) (NGRANCellReportList, int, error) {
	var v NGRANCellReportList
	err := ngapcodec.DecodeParameter("NGRAN-CellReportList", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANCellReportList{}, 0, err
	}
	return v, len(data), err
}
