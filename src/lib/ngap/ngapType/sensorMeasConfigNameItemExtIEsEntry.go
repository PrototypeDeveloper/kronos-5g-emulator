// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SensorMeasConfigNameItemExtIEsEntryPresent uint32

const (
	SensorMeasConfigNameItemExtIEsEntryPresentNothing SensorMeasConfigNameItemExtIEsEntryPresent = iota
	SensorMeasConfigNameItemExtIEsEntryPresentUnknown
)

type SensorMeasConfigNameItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SensorMeasConfigNameItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SensorMeasConfigNameItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SensorMeasConfigNameItem-ExtIEsEntry", v)
}

// DecodeSensorMeasConfigNameItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSensorMeasConfigNameItemExtIEsEntry(data []byte) (SensorMeasConfigNameItemExtIEsEntry, int, error) {
	var v SensorMeasConfigNameItemExtIEsEntry
	err := ngapcodec.DecodeParameter("SensorMeasConfigNameItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SensorMeasConfigNameItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
