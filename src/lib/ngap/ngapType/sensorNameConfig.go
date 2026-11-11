// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SensorNameConfig represents ASN.1 SensorNameConfig.
type SensorNameConfigPresent uint32

const (
	SensorNameConfigPresentNothing SensorNameConfigPresent = iota
	SensorNameConfigPresentUncompensatedBarometricConfig
	SensorNameConfigPresentUeSpeedConfig
	SensorNameConfigPresentUeOrientationConfig
	SensorNameConfigPresentChoiceExtensions
)

type SensorNameConfig struct {
	Present                       SensorNameConfigPresent
	UncompensatedBarometricConfig SensorNameConfigUncompensatedBarometricConfig `aper:"valueExt,valueLB:0,valueUB:0"`
	UeSpeedConfig                 SensorNameConfigUeSpeedConfig                 `aper:"valueExt,valueLB:0,valueUB:0"`
	UeOrientationConfig           SensorNameConfigUeOrientationConfig           `aper:"valueExt,valueLB:0,valueUB:0"`
	ChoiceExtensions              SensorNameConfigExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SensorNameConfig) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SensorNameConfig", v)
}

// DecodeSensorNameConfig decodes a standalone APER value and reports the consumed octets.
func DecodeSensorNameConfig(data []byte) (SensorNameConfig, int, error) {
	var v SensorNameConfig
	err := ngapcodec.DecodeParameter("SensorNameConfig", data, &v)
	if ngapcodec.IsFatal(err) {
		return SensorNameConfig{}, 0, err
	}
	return v, len(data), err
}
