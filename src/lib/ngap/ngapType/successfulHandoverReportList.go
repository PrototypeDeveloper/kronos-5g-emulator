// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SuccessfulHandoverReportList represents ASN.1 SuccessfulHandoverReportList.
type SuccessfulHandoverReportList struct {
	List []SuccessfulHandoverReportItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SuccessfulHandoverReportList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SuccessfulHandoverReportList", v)
}

// DecodeSuccessfulHandoverReportList decodes a standalone APER value and reports the consumed octets.
func DecodeSuccessfulHandoverReportList(data []byte) (SuccessfulHandoverReportList, int, error) {
	var v SuccessfulHandoverReportList
	err := ngapcodec.DecodeParameter("SuccessfulHandoverReportList", data, &v)
	if ngapcodec.IsFatal(err) {
		return SuccessfulHandoverReportList{}, 0, err
	}
	return v, len(data), err
}
