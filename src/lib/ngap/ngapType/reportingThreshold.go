// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ReportingThreshold represents ASN.1 ReportingThreshold.
type ReportingThreshold struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:4000000000"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ReportingThreshold) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ReportingThreshold", v)
}

// DecodeReportingThreshold decodes a standalone APER value and reports the consumed octets.
func DecodeReportingThreshold(data []byte) (ReportingThreshold, int, error) {
	var v ReportingThreshold
	err := ngapcodec.DecodeParameter("ReportingThreshold", data, &v)
	if ngapcodec.IsFatal(err) {
		return ReportingThreshold{}, 0, err
	}
	return v, len(data), err
}
