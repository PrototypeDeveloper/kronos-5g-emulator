// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BeamMeasurementsReportQuantityRSRP represents ASN.1 BeamMeasurementsReportQuantityRSRP.
const (
	BeamMeasurementsReportQuantityRSRPPresentTrue  Enumerated = 0
	BeamMeasurementsReportQuantityRSRPPresentFalse Enumerated = 1
)

type BeamMeasurementsReportQuantityRSRP struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v BeamMeasurementsReportQuantityRSRP) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:BeamMeasurementsReportQuantityRSRP", v)
}

// DecodeBeamMeasurementsReportQuantityRSRP decodes a standalone APER value and reports consumed octets.
func DecodeBeamMeasurementsReportQuantityRSRP(data []byte) (BeamMeasurementsReportQuantityRSRP, int, error) {
	var v BeamMeasurementsReportQuantityRSRP
	err := ngapcodec.DecodeParameter("inline:BeamMeasurementsReportQuantityRSRP", data, &v)
	if ngapcodec.IsFatal(err) {
		return BeamMeasurementsReportQuantityRSRP{}, 0, err
	}
	return v, len(data), err
}
