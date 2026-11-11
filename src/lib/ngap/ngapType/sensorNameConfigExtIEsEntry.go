// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SensorNameConfigExtIEsEntryPresent uint32

const (
	SensorNameConfigExtIEsEntryPresentNothing SensorNameConfigExtIEsEntryPresent = iota
	SensorNameConfigExtIEsEntryPresentUnknown
)

type SensorNameConfigExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      SensorNameConfigExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SensorNameConfigExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SensorNameConfig-ExtIEsEntry", v)
}

// DecodeSensorNameConfigExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSensorNameConfigExtIEsEntry(data []byte) (SensorNameConfigExtIEsEntry, int, error) {
	var v SensorNameConfigExtIEsEntry
	err := ngapcodec.DecodeParameter("SensorNameConfig-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SensorNameConfigExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
