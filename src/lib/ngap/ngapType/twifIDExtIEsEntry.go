// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TWIFIDExtIEsEntryPresent uint32

const (
	TWIFIDExtIEsEntryPresentNothing TWIFIDExtIEsEntryPresent = iota
	TWIFIDExtIEsEntryPresentUnknown
)

type TWIFIDExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      TWIFIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TWIFIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TWIF-ID-ExtIEsEntry", v)
}

// DecodeTWIFIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTWIFIDExtIEsEntry(data []byte) (TWIFIDExtIEsEntry, int, error) {
	var v TWIFIDExtIEsEntry
	err := ngapcodec.DecodeParameter("TWIF-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TWIFIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
