// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type BluetoothMeasurementConfigurationExtIEsEntryPresent uint32

const (
	BluetoothMeasurementConfigurationExtIEsEntryPresentNothing BluetoothMeasurementConfigurationExtIEsEntryPresent = iota
	BluetoothMeasurementConfigurationExtIEsEntryPresentUnknown
)

type BluetoothMeasurementConfigurationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      BluetoothMeasurementConfigurationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BluetoothMeasurementConfigurationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BluetoothMeasurementConfiguration-ExtIEsEntry", v)
}

// DecodeBluetoothMeasurementConfigurationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeBluetoothMeasurementConfigurationExtIEsEntry(data []byte) (BluetoothMeasurementConfigurationExtIEsEntry, int, error) {
	var v BluetoothMeasurementConfigurationExtIEsEntry
	err := ngapcodec.DecodeParameter("BluetoothMeasurementConfiguration-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return BluetoothMeasurementConfigurationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
