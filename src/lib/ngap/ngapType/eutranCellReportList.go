// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EUTRANCellReportList represents ASN.1 EUTRAN-CellReportList.
type EUTRANCellReportList struct {
	List []EUTRANCellReportItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRANCellReportList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRAN-CellReportList", v)
}

// DecodeEUTRANCellReportList decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRANCellReportList(data []byte) (EUTRANCellReportList, int, error) {
	var v EUTRANCellReportList
	err := ngapcodec.DecodeParameter("EUTRAN-CellReportList", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRANCellReportList{}, 0, err
	}
	return v, len(data), err
}
