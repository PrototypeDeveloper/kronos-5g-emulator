// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BluetoothName represents ASN.1 BluetoothName.
type BluetoothName struct {
	Value []byte `aper:"sizeLB:1,sizeUB:248"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BluetoothName) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("BluetoothName", v) }

// DecodeBluetoothName decodes a standalone APER value and reports the consumed octets.
func DecodeBluetoothName(data []byte) (BluetoothName, int, error) {
	var v BluetoothName
	err := ngapcodec.DecodeParameter("BluetoothName", data, &v)
	if ngapcodec.IsFatal(err) {
		return BluetoothName{}, 0, err
	}
	return v, len(data), err
}
