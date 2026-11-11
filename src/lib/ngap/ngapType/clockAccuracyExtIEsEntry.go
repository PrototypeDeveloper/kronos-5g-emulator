// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ClockAccuracyExtIEsEntryPresent uint32

const (
	ClockAccuracyExtIEsEntryPresentNothing ClockAccuracyExtIEsEntryPresent = iota
	ClockAccuracyExtIEsEntryPresentUnknown
)

type ClockAccuracyExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      ClockAccuracyExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ClockAccuracyExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ClockAccuracy-ExtIEsEntry", v)
}

// DecodeClockAccuracyExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeClockAccuracyExtIEsEntry(data []byte) (ClockAccuracyExtIEsEntry, int, error) {
	var v ClockAccuracyExtIEsEntry
	err := ngapcodec.DecodeParameter("ClockAccuracy-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ClockAccuracyExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
