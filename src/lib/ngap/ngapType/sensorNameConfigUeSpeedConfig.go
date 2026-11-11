// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SensorNameConfigUeSpeedConfig represents ASN.1 SensorNameConfigUeSpeedConfig.
const (
	SensorNameConfigUeSpeedConfigPresentTrue Enumerated = 0
)

type SensorNameConfigUeSpeedConfig struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v SensorNameConfigUeSpeedConfig) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:SensorNameConfigUeSpeedConfig", v)
}

// DecodeSensorNameConfigUeSpeedConfig decodes a standalone APER value and reports consumed octets.
func DecodeSensorNameConfigUeSpeedConfig(data []byte) (SensorNameConfigUeSpeedConfig, int, error) {
	var v SensorNameConfigUeSpeedConfig
	err := ngapcodec.DecodeParameter("inline:SensorNameConfigUeSpeedConfig", data, &v)
	if ngapcodec.IsFatal(err) {
		return SensorNameConfigUeSpeedConfig{}, 0, err
	}
	return v, len(data), err
}
