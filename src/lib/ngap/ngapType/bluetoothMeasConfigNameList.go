// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BluetoothMeasConfigNameList represents ASN.1 BluetoothMeasConfigNameList.
type BluetoothMeasConfigNameList struct {
	List []BluetoothMeasConfigNameItem `aper:"sizeLB:1,sizeUB:4"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BluetoothMeasConfigNameList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BluetoothMeasConfigNameList", v)
}

// DecodeBluetoothMeasConfigNameList decodes a standalone APER value and reports the consumed octets.
func DecodeBluetoothMeasConfigNameList(data []byte) (BluetoothMeasConfigNameList, int, error) {
	var v BluetoothMeasConfigNameList
	err := ngapcodec.DecodeParameter("BluetoothMeasConfigNameList", data, &v)
	if ngapcodec.IsFatal(err) {
		return BluetoothMeasConfigNameList{}, 0, err
	}
	return v, len(data), err
}
