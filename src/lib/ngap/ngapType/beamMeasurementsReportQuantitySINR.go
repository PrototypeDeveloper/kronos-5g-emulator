// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BeamMeasurementsReportQuantitySINR represents ASN.1 BeamMeasurementsReportQuantitySINR.
const (
	BeamMeasurementsReportQuantitySINRPresentTrue  Enumerated = 0
	BeamMeasurementsReportQuantitySINRPresentFalse Enumerated = 1
)

type BeamMeasurementsReportQuantitySINR struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v BeamMeasurementsReportQuantitySINR) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:BeamMeasurementsReportQuantitySINR", v)
}

// DecodeBeamMeasurementsReportQuantitySINR decodes a standalone APER value and reports consumed octets.
func DecodeBeamMeasurementsReportQuantitySINR(data []byte) (BeamMeasurementsReportQuantitySINR, int, error) {
	var v BeamMeasurementsReportQuantitySINR
	err := ngapcodec.DecodeParameter("inline:BeamMeasurementsReportQuantitySINR", data, &v)
	if ngapcodec.IsFatal(err) {
		return BeamMeasurementsReportQuantitySINR{}, 0, err
	}
	return v, len(data), err
}
