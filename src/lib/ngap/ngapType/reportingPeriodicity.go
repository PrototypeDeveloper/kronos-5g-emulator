// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ReportingPeriodicity represents ASN.1 ReportingPeriodicity.
const (
	ReportingPeriodicityPresentStop    Enumerated = 0
	ReportingPeriodicityPresentSingle  Enumerated = 1
	ReportingPeriodicityPresentMs1000  Enumerated = 2
	ReportingPeriodicityPresentMs2000  Enumerated = 3
	ReportingPeriodicityPresentMs5000  Enumerated = 4
	ReportingPeriodicityPresentMs10000 Enumerated = 5
)

type ReportingPeriodicity struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:5"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ReportingPeriodicity) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ReportingPeriodicity", v)
}

// DecodeReportingPeriodicity decodes a standalone APER value and reports the consumed octets.
func DecodeReportingPeriodicity(data []byte) (ReportingPeriodicity, int, error) {
	var v ReportingPeriodicity
	err := ngapcodec.DecodeParameter("ReportingPeriodicity", data, &v)
	if ngapcodec.IsFatal(err) {
		return ReportingPeriodicity{}, 0, err
	}
	return v, len(data), err
}
