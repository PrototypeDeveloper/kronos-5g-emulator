// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AerialUEReportingPeriodicity represents ASN.1 AerialUEReportingPeriodicity.
const (
	AerialUEReportingPeriodicityPresentMs120   Enumerated = 0
	AerialUEReportingPeriodicityPresentMs240   Enumerated = 1
	AerialUEReportingPeriodicityPresentMs480   Enumerated = 2
	AerialUEReportingPeriodicityPresentMs640   Enumerated = 3
	AerialUEReportingPeriodicityPresentMs1024  Enumerated = 4
	AerialUEReportingPeriodicityPresentMs2048  Enumerated = 5
	AerialUEReportingPeriodicityPresentMs5120  Enumerated = 6
	AerialUEReportingPeriodicityPresentMs10240 Enumerated = 7
	AerialUEReportingPeriodicityPresentMs20480 Enumerated = 8
	AerialUEReportingPeriodicityPresentMs40960 Enumerated = 9
	AerialUEReportingPeriodicityPresentMin1    Enumerated = 10
	AerialUEReportingPeriodicityPresentMin6    Enumerated = 11
	AerialUEReportingPeriodicityPresentMin12   Enumerated = 12
	AerialUEReportingPeriodicityPresentMin30   Enumerated = 13
)

type AerialUEReportingPeriodicity struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:13"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AerialUEReportingPeriodicity) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AerialUEReportingPeriodicity", v)
}

// DecodeAerialUEReportingPeriodicity decodes a standalone APER value and reports the consumed octets.
func DecodeAerialUEReportingPeriodicity(data []byte) (AerialUEReportingPeriodicity, int, error) {
	var v AerialUEReportingPeriodicity
	err := ngapcodec.DecodeParameter("AerialUEReportingPeriodicity", data, &v)
	if ngapcodec.IsFatal(err) {
		return AerialUEReportingPeriodicity{}, 0, err
	}
	return v, len(data), err
}
