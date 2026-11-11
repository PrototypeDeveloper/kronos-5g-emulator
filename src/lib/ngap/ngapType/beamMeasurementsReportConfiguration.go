// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BeamMeasurementsReportConfiguration represents ASN.1 BeamMeasurementsReportConfiguration.
type BeamMeasurementsReportConfiguration struct {
	HasBeamMeasurementsReportQuantity bool
	BeamMeasurementsReportQuantity    BeamMeasurementsReportQuantity
	HasMaxNrofRSIndexesToReport       bool
	MaxNrofRSIndexesToReport          MaxNrofRSIndexesToReport
	HasIEExtensions                   bool
	IEExtensions                      []BeamMeasurementsReportConfigurationExtIEsEntry
	UnknownExtensions                 []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BeamMeasurementsReportConfiguration) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BeamMeasurementsReportConfiguration", v)
}

// DecodeBeamMeasurementsReportConfiguration decodes a standalone APER value and reports the consumed octets.
func DecodeBeamMeasurementsReportConfiguration(data []byte) (BeamMeasurementsReportConfiguration, int, error) {
	var v BeamMeasurementsReportConfiguration
	err := ngapcodec.DecodeParameter("BeamMeasurementsReportConfiguration", data, &v)
	if ngapcodec.IsFatal(err) {
		return BeamMeasurementsReportConfiguration{}, 0, err
	}
	return v, len(data), err
}
