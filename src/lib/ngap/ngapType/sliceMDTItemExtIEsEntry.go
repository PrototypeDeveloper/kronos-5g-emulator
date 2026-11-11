// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SliceMDTItemExtIEsEntryPresent uint32

const (
	SliceMDTItemExtIEsEntryPresentNothing SliceMDTItemExtIEsEntryPresent = iota
	SliceMDTItemExtIEsEntryPresentUnknown
)

type SliceMDTItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SliceMDTItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SliceMDTItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SliceMDTItem-ExtIEsEntry", v)
}

// DecodeSliceMDTItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSliceMDTItemExtIEsEntry(data []byte) (SliceMDTItemExtIEsEntry, int, error) {
	var v SliceMDTItemExtIEsEntry
	err := ngapcodec.DecodeParameter("SliceMDTItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SliceMDTItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
