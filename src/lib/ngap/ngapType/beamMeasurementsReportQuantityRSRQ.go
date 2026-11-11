// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BeamMeasurementsReportQuantityRSRQ represents ASN.1 BeamMeasurementsReportQuantityRSRQ.
const (
	BeamMeasurementsReportQuantityRSRQPresentTrue  Enumerated = 0
	BeamMeasurementsReportQuantityRSRQPresentFalse Enumerated = 1
)

type BeamMeasurementsReportQuantityRSRQ struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v BeamMeasurementsReportQuantityRSRQ) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:BeamMeasurementsReportQuantityRSRQ", v)
}

// DecodeBeamMeasurementsReportQuantityRSRQ decodes a standalone APER value and reports consumed octets.
func DecodeBeamMeasurementsReportQuantityRSRQ(data []byte) (BeamMeasurementsReportQuantityRSRQ, int, error) {
	var v BeamMeasurementsReportQuantityRSRQ
	err := ngapcodec.DecodeParameter("inline:BeamMeasurementsReportQuantityRSRQ", data, &v)
	if ngapcodec.IsFatal(err) {
		return BeamMeasurementsReportQuantityRSRQ{}, 0, err
	}
	return v, len(data), err
}
