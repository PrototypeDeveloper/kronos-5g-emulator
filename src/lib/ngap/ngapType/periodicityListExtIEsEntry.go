// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PeriodicityListExtIEsEntryPresent uint32

const (
	PeriodicityListExtIEsEntryPresentNothing PeriodicityListExtIEsEntryPresent = iota
	PeriodicityListExtIEsEntryPresentUnknown
)

type PeriodicityListExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PeriodicityListExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PeriodicityListExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PeriodicityList-ExtIEsEntry", v)
}

// DecodePeriodicityListExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePeriodicityListExtIEsEntry(data []byte) (PeriodicityListExtIEsEntry, int, error) {
	var v PeriodicityListExtIEsEntry
	err := ngapcodec.DecodeParameter("PeriodicityList-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PeriodicityListExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
