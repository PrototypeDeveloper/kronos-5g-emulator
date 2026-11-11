// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type GUAMIListItemExtIEsEntryPresent uint32

const (
	GUAMIListItemExtIEsEntryPresentNothing GUAMIListItemExtIEsEntryPresent = iota
	GUAMIListItemExtIEsEntryPresentUnknown
)

type GUAMIListItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      GUAMIListItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GUAMIListItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GUAMIList-Item-ExtIEsEntry", v)
}

// DecodeGUAMIListItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeGUAMIListItemExtIEsEntry(data []byte) (GUAMIListItemExtIEsEntry, int, error) {
	var v GUAMIListItemExtIEsEntry
	err := ngapcodec.DecodeParameter("GUAMIList-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return GUAMIListItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
