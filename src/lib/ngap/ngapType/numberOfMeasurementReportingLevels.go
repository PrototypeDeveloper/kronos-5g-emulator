// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NumberOfMeasurementReportingLevels represents ASN.1 NumberOfMeasurementReportingLevels.
const (
	NumberOfMeasurementReportingLevelsPresentN2  Enumerated = 0
	NumberOfMeasurementReportingLevelsPresentN3  Enumerated = 1
	NumberOfMeasurementReportingLevelsPresentN4  Enumerated = 2
	NumberOfMeasurementReportingLevelsPresentN5  Enumerated = 3
	NumberOfMeasurementReportingLevelsPresentN10 Enumerated = 4
	NumberOfMeasurementReportingLevelsPresentN0  Enumerated = 5
)

type NumberOfMeasurementReportingLevels struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:4"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NumberOfMeasurementReportingLevels) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NumberOfMeasurementReportingLevels", v)
}

// DecodeNumberOfMeasurementReportingLevels decodes a standalone APER value and reports the consumed octets.
func DecodeNumberOfMeasurementReportingLevels(data []byte) (NumberOfMeasurementReportingLevels, int, error) {
	var v NumberOfMeasurementReportingLevels
	err := ngapcodec.DecodeParameter("NumberOfMeasurementReportingLevels", data, &v)
	if ngapcodec.IsFatal(err) {
		return NumberOfMeasurementReportingLevels{}, 0, err
	}
	return v, len(data), err
}
