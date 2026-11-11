// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MDTModeNrExtIEsEntryPresent uint32

const (
	MDTModeNrExtIEsEntryPresentNothing MDTModeNrExtIEsEntryPresent = iota
	MDTModeNrExtIEsEntryPresentUnknown
)

type MDTModeNrExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      MDTModeNrExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTModeNrExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MDTModeNr-ExtIEsEntry", v)
}

// DecodeMDTModeNrExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMDTModeNrExtIEsEntry(data []byte) (MDTModeNrExtIEsEntry, int, error) {
	var v MDTModeNrExtIEsEntry
	err := ngapcodec.DecodeParameter("MDTModeNr-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTModeNrExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
