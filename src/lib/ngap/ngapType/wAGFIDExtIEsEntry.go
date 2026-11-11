// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type WAGFIDExtIEsEntryPresent uint32

const (
	WAGFIDExtIEsEntryPresentNothing WAGFIDExtIEsEntryPresent = iota
	WAGFIDExtIEsEntryPresentUnknown
)

type WAGFIDExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      WAGFIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v WAGFIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("W-AGF-ID-ExtIEsEntry", v)
}

// DecodeWAGFIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeWAGFIDExtIEsEntry(data []byte) (WAGFIDExtIEsEntry, int, error) {
	var v WAGFIDExtIEsEntry
	err := ngapcodec.DecodeParameter("W-AGF-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return WAGFIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
