// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PeriodicityBoundExtIEsEntryPresent uint32

const (
	PeriodicityBoundExtIEsEntryPresentNothing PeriodicityBoundExtIEsEntryPresent = iota
	PeriodicityBoundExtIEsEntryPresentUnknown
)

type PeriodicityBoundExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PeriodicityBoundExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PeriodicityBoundExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PeriodicityBound-ExtIEsEntry", v)
}

// DecodePeriodicityBoundExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePeriodicityBoundExtIEsEntry(data []byte) (PeriodicityBoundExtIEsEntry, int, error) {
	var v PeriodicityBoundExtIEsEntry
	err := ngapcodec.DecodeParameter("PeriodicityBound-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PeriodicityBoundExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
