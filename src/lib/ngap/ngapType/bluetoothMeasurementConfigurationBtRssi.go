// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BluetoothMeasurementConfigurationBtRssi represents ASN.1 BluetoothMeasurementConfigurationBtRssi.
const (
	BluetoothMeasurementConfigurationBtRssiPresentTrue Enumerated = 0
)

type BluetoothMeasurementConfigurationBtRssi struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v BluetoothMeasurementConfigurationBtRssi) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:BluetoothMeasurementConfigurationBtRssi", v)
}

// DecodeBluetoothMeasurementConfigurationBtRssi decodes a standalone APER value and reports consumed octets.
func DecodeBluetoothMeasurementConfigurationBtRssi(data []byte) (BluetoothMeasurementConfigurationBtRssi, int, error) {
	var v BluetoothMeasurementConfigurationBtRssi
	err := ngapcodec.DecodeParameter("inline:BluetoothMeasurementConfigurationBtRssi", data, &v)
	if ngapcodec.IsFatal(err) {
		return BluetoothMeasurementConfigurationBtRssi{}, 0, err
	}
	return v, len(data), err
}
