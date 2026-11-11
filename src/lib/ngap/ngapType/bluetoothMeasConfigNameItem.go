// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BluetoothMeasConfigNameItem represents ASN.1 BluetoothMeasConfigNameItem.
type BluetoothMeasConfigNameItem struct {
	BluetoothName     BluetoothName
	HasIEExtensions   bool
	IEExtensions      []BluetoothMeasConfigNameItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BluetoothMeasConfigNameItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BluetoothMeasConfigNameItem", v)
}

// DecodeBluetoothMeasConfigNameItem decodes a standalone APER value and reports the consumed octets.
func DecodeBluetoothMeasConfigNameItem(data []byte) (BluetoothMeasConfigNameItem, int, error) {
	var v BluetoothMeasConfigNameItem
	err := ngapcodec.DecodeParameter("BluetoothMeasConfigNameItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return BluetoothMeasConfigNameItem{}, 0, err
	}
	return v, len(data), err
}
