// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MDTLocationInfoExtIEsEntryPresent uint32

const (
	MDTLocationInfoExtIEsEntryPresentNothing MDTLocationInfoExtIEsEntryPresent = iota
	MDTLocationInfoExtIEsEntryPresentUnknown
)

type MDTLocationInfoExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MDTLocationInfoExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTLocationInfoExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MDT-Location-Info-ExtIEsEntry", v)
}

// DecodeMDTLocationInfoExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMDTLocationInfoExtIEsEntry(data []byte) (MDTLocationInfoExtIEsEntry, int, error) {
	var v MDTLocationInfoExtIEsEntry
	err := ngapcodec.DecodeParameter("MDT-Location-Info-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTLocationInfoExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
