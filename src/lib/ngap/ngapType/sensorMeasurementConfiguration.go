// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SensorMeasurementConfiguration represents ASN.1 SensorMeasurementConfiguration.
type SensorMeasurementConfiguration struct {
	SensorMeasConfig            SensorMeasConfig
	HasSensorMeasConfigNameList bool
	SensorMeasConfigNameList    SensorMeasConfigNameList
	HasIEExtensions             bool
	IEExtensions                []SensorMeasurementConfigurationExtIEsEntry
	UnknownExtensions           []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SensorMeasurementConfiguration) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SensorMeasurementConfiguration", v)
}

// DecodeSensorMeasurementConfiguration decodes a standalone APER value and reports the consumed octets.
func DecodeSensorMeasurementConfiguration(data []byte) (SensorMeasurementConfiguration, int, error) {
	var v SensorMeasurementConfiguration
	err := ngapcodec.DecodeParameter("SensorMeasurementConfiguration", data, &v)
	if ngapcodec.IsFatal(err) {
		return SensorMeasurementConfiguration{}, 0, err
	}
	return v, len(data), err
}
