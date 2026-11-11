// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type BluetoothMeasConfigNameItemExtIEsEntryPresent uint32

const (
	BluetoothMeasConfigNameItemExtIEsEntryPresentNothing BluetoothMeasConfigNameItemExtIEsEntryPresent = iota
	BluetoothMeasConfigNameItemExtIEsEntryPresentUnknown
)

type BluetoothMeasConfigNameItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      BluetoothMeasConfigNameItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BluetoothMeasConfigNameItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BluetoothMeasConfigNameItem-ExtIEsEntry", v)
}

// DecodeBluetoothMeasConfigNameItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeBluetoothMeasConfigNameItemExtIEsEntry(data []byte) (BluetoothMeasConfigNameItemExtIEsEntry, int, error) {
	var v BluetoothMeasConfigNameItemExtIEsEntry
	err := ngapcodec.DecodeParameter("BluetoothMeasConfigNameItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return BluetoothMeasConfigNameItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
