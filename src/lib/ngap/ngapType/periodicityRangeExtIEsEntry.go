// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PeriodicityRangeExtIEsEntryPresent uint32

const (
	PeriodicityRangeExtIEsEntryPresentNothing PeriodicityRangeExtIEsEntryPresent = iota
	PeriodicityRangeExtIEsEntryPresentUnknown
)

type PeriodicityRangeExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      PeriodicityRangeExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PeriodicityRangeExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PeriodicityRange-ExtIEsEntry", v)
}

// DecodePeriodicityRangeExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePeriodicityRangeExtIEsEntry(data []byte) (PeriodicityRangeExtIEsEntry, int, error) {
	var v PeriodicityRangeExtIEsEntry
	err := ngapcodec.DecodeParameter("PeriodicityRange-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PeriodicityRangeExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
