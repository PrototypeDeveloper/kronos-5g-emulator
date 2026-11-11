// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type N3IWFIDExtIEsEntryPresent uint32

const (
	N3IWFIDExtIEsEntryPresentNothing N3IWFIDExtIEsEntryPresent = iota
	N3IWFIDExtIEsEntryPresentUnknown
)

type N3IWFIDExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      N3IWFIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v N3IWFIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("N3IWF-ID-ExtIEsEntry", v)
}

// DecodeN3IWFIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeN3IWFIDExtIEsEntry(data []byte) (N3IWFIDExtIEsEntry, int, error) {
	var v N3IWFIDExtIEsEntry
	err := ngapcodec.DecodeParameter("N3IWF-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return N3IWFIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
