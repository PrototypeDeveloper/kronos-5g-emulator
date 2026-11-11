// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SuccessfulPSCellChangeReportList represents ASN.1 SuccessfulPSCellChangeReportList.
type SuccessfulPSCellChangeReportList struct {
	List []SuccessfulPSCellChangeReportItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SuccessfulPSCellChangeReportList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SuccessfulPSCellChangeReportList", v)
}

// DecodeSuccessfulPSCellChangeReportList decodes a standalone APER value and reports the consumed octets.
func DecodeSuccessfulPSCellChangeReportList(data []byte) (SuccessfulPSCellChangeReportList, int, error) {
	var v SuccessfulPSCellChangeReportList
	err := ngapcodec.DecodeParameter("SuccessfulPSCellChangeReportList", data, &v)
	if ngapcodec.IsFatal(err) {
		return SuccessfulPSCellChangeReportList{}, 0, err
	}
	return v, len(data), err
}
