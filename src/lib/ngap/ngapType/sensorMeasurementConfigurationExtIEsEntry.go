// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SensorMeasurementConfigurationExtIEsEntryPresent uint32

const (
	SensorMeasurementConfigurationExtIEsEntryPresentNothing SensorMeasurementConfigurationExtIEsEntryPresent = iota
	SensorMeasurementConfigurationExtIEsEntryPresentUnknown
)

type SensorMeasurementConfigurationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SensorMeasurementConfigurationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SensorMeasurementConfigurationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SensorMeasurementConfiguration-ExtIEsEntry", v)
}

// DecodeSensorMeasurementConfigurationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSensorMeasurementConfigurationExtIEsEntry(data []byte) (SensorMeasurementConfigurationExtIEsEntry, int, error) {
	var v SensorMeasurementConfigurationExtIEsEntry
	err := ngapcodec.DecodeParameter("SensorMeasurementConfiguration-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SensorMeasurementConfigurationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
