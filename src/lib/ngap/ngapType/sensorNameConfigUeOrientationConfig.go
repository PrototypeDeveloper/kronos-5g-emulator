// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SensorNameConfigUeOrientationConfig represents ASN.1 SensorNameConfigUeOrientationConfig.
const (
	SensorNameConfigUeOrientationConfigPresentTrue Enumerated = 0
)

type SensorNameConfigUeOrientationConfig struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v SensorNameConfigUeOrientationConfig) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:SensorNameConfigUeOrientationConfig", v)
}

// DecodeSensorNameConfigUeOrientationConfig decodes a standalone APER value and reports consumed octets.
func DecodeSensorNameConfigUeOrientationConfig(data []byte) (SensorNameConfigUeOrientationConfig, int, error) {
	var v SensorNameConfigUeOrientationConfig
	err := ngapcodec.DecodeParameter("inline:SensorNameConfigUeOrientationConfig", data, &v)
	if ngapcodec.IsFatal(err) {
		return SensorNameConfigUeOrientationConfig{}, 0, err
	}
	return v, len(data), err
}
