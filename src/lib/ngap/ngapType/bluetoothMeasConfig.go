// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BluetoothMeasConfig represents ASN.1 BluetoothMeasConfig.
const (
	BluetoothMeasConfigPresentSetup Enumerated = 0
)

type BluetoothMeasConfig struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BluetoothMeasConfig) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BluetoothMeasConfig", v)
}

// DecodeBluetoothMeasConfig decodes a standalone APER value and reports the consumed octets.
func DecodeBluetoothMeasConfig(data []byte) (BluetoothMeasConfig, int, error) {
	var v BluetoothMeasConfig
	err := ngapcodec.DecodeParameter("BluetoothMeasConfig", data, &v)
	if ngapcodec.IsFatal(err) {
		return BluetoothMeasConfig{}, 0, err
	}
	return v, len(data), err
}
