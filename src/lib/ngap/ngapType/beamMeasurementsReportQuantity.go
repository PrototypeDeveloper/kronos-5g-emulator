// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BeamMeasurementsReportQuantity represents ASN.1 BeamMeasurementsReportQuantity.
type BeamMeasurementsReportQuantity struct {
	RSRP              BeamMeasurementsReportQuantityRSRP `aper:"valueExt,valueLB:0,valueUB:0"`
	RSRQ              BeamMeasurementsReportQuantityRSRQ `aper:"valueExt,valueLB:0,valueUB:0"`
	SINR              BeamMeasurementsReportQuantitySINR `aper:"valueExt,valueLB:0,valueUB:0"`
	HasIEExtensions   bool
	IEExtensions      []BeamMeasurementsReportQuantityExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BeamMeasurementsReportQuantity) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BeamMeasurementsReportQuantity", v)
}

// DecodeBeamMeasurementsReportQuantity decodes a standalone APER value and reports the consumed octets.
func DecodeBeamMeasurementsReportQuantity(data []byte) (BeamMeasurementsReportQuantity, int, error) {
	var v BeamMeasurementsReportQuantity
	err := ngapcodec.DecodeParameter("BeamMeasurementsReportQuantity", data, &v)
	if ngapcodec.IsFatal(err) {
		return BeamMeasurementsReportQuantity{}, 0, err
	}
	return v, len(data), err
}
