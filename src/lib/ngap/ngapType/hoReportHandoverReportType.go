// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// HOReportHandoverReportType represents ASN.1 HOReportHandoverReportType.
const (
	HOReportHandoverReportTypePresentHoTooEarly          Enumerated = 0
	HOReportHandoverReportTypePresentHoToWrongCell       Enumerated = 1
	HOReportHandoverReportTypePresentIntersystemPingPong Enumerated = 2
)

type HOReportHandoverReportType struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value.
func (v HOReportHandoverReportType) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:HOReportHandoverReportType", v)
}

// DecodeHOReportHandoverReportType decodes a standalone APER value and reports consumed octets.
func DecodeHOReportHandoverReportType(data []byte) (HOReportHandoverReportType, int, error) {
	var v HOReportHandoverReportType
	err := ngapcodec.DecodeParameter("inline:HOReportHandoverReportType", data, &v)
	if ngapcodec.IsFatal(err) {
		return HOReportHandoverReportType{}, 0, err
	}
	return v, len(data), err
}
