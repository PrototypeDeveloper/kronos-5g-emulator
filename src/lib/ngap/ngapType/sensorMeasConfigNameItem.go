// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SensorMeasConfigNameItem represents ASN.1 SensorMeasConfigNameItem.
type SensorMeasConfigNameItem struct {
	SensorNameConfig  SensorNameConfig
	HasIEExtensions   bool
	IEExtensions      []SensorMeasConfigNameItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SensorMeasConfigNameItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SensorMeasConfigNameItem", v)
}

// DecodeSensorMeasConfigNameItem decodes a standalone APER value and reports the consumed octets.
func DecodeSensorMeasConfigNameItem(data []byte) (SensorMeasConfigNameItem, int, error) {
	var v SensorMeasConfigNameItem
	err := ngapcodec.DecodeParameter("SensorMeasConfigNameItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return SensorMeasConfigNameItem{}, 0, err
	}
	return v, len(data), err
}
