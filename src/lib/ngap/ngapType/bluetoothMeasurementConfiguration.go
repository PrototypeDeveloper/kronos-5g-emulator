// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BluetoothMeasurementConfiguration represents ASN.1 BluetoothMeasurementConfiguration.
type BluetoothMeasurementConfiguration struct {
	BluetoothMeasConfig            BluetoothMeasConfig
	HasBluetoothMeasConfigNameList bool
	BluetoothMeasConfigNameList    BluetoothMeasConfigNameList
	HasBtRssi                      bool
	BtRssi                         BluetoothMeasurementConfigurationBtRssi `aper:"valueExt,valueLB:0,valueUB:0"`
	HasIEExtensions                bool
	IEExtensions                   []BluetoothMeasurementConfigurationExtIEsEntry
	UnknownExtensions              []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BluetoothMeasurementConfiguration) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BluetoothMeasurementConfiguration", v)
}

// DecodeBluetoothMeasurementConfiguration decodes a standalone APER value and reports the consumed octets.
func DecodeBluetoothMeasurementConfiguration(data []byte) (BluetoothMeasurementConfiguration, int, error) {
	var v BluetoothMeasurementConfiguration
	err := ngapcodec.DecodeParameter("BluetoothMeasurementConfiguration", data, &v)
	if ngapcodec.IsFatal(err) {
		return BluetoothMeasurementConfiguration{}, 0, err
	}
	return v, len(data), err
}
