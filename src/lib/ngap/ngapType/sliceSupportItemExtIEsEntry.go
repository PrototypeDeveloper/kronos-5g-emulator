// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SliceSupportItemExtIEsEntryPresent uint32

const (
	SliceSupportItemExtIEsEntryPresentNothing SliceSupportItemExtIEsEntryPresent = iota
	SliceSupportItemExtIEsEntryPresentUnknown
)

type SliceSupportItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SliceSupportItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SliceSupportItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SliceSupportItem-ExtIEsEntry", v)
}

// DecodeSliceSupportItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSliceSupportItemExtIEsEntry(data []byte) (SliceSupportItemExtIEsEntry, int, error) {
	var v SliceSupportItemExtIEsEntry
	err := ngapcodec.DecodeParameter("SliceSupportItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SliceSupportItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
