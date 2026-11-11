// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SensorNameConfigUncompensatedBarometricConfig represents ASN.1 SensorNameConfigUncompensatedBarometricConfig.
const (
	SensorNameConfigUncompensatedBarometricConfigPresentTrue Enumerated = 0
)

type SensorNameConfigUncompensatedBarometricConfig struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v SensorNameConfigUncompensatedBarometricConfig) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:SensorNameConfigUncompensatedBarometricConfig", v)
}

// DecodeSensorNameConfigUncompensatedBarometricConfig decodes a standalone APER value and reports consumed octets.
func DecodeSensorNameConfigUncompensatedBarometricConfig(data []byte) (SensorNameConfigUncompensatedBarometricConfig, int, error) {
	var v SensorNameConfigUncompensatedBarometricConfig
	err := ngapcodec.DecodeParameter("inline:SensorNameConfigUncompensatedBarometricConfig", data, &v)
	if ngapcodec.IsFatal(err) {
		return SensorNameConfigUncompensatedBarometricConfig{}, 0, err
	}
	return v, len(data), err
}
