// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SensorMeasConfig represents ASN.1 SensorMeasConfig.
const (
	SensorMeasConfigPresentSetup Enumerated = 0
)

type SensorMeasConfig struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SensorMeasConfig) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SensorMeasConfig", v)
}

// DecodeSensorMeasConfig decodes a standalone APER value and reports the consumed octets.
func DecodeSensorMeasConfig(data []byte) (SensorMeasConfig, int, error) {
	var v SensorMeasConfig
	err := ngapcodec.DecodeParameter("SensorMeasConfig", data, &v)
	if ngapcodec.IsFatal(err) {
		return SensorMeasConfig{}, 0, err
	}
	return v, len(data), err
}
