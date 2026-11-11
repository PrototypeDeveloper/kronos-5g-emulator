// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EarlyMeasurement represents ASN.1 EarlyMeasurement.
const (
	EarlyMeasurementPresentTrue Enumerated = 0
)

type EarlyMeasurement struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EarlyMeasurement) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EarlyMeasurement", v)
}

// DecodeEarlyMeasurement decodes a standalone APER value and reports the consumed octets.
func DecodeEarlyMeasurement(data []byte) (EarlyMeasurement, int, error) {
	var v EarlyMeasurement
	err := ngapcodec.DecodeParameter("EarlyMeasurement", data, &v)
	if ngapcodec.IsFatal(err) {
		return EarlyMeasurement{}, 0, err
	}
	return v, len(data), err
}
