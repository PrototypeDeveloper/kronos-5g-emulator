// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SliceSupportQMCItemExtIEsEntryPresent uint32

const (
	SliceSupportQMCItemExtIEsEntryPresentNothing SliceSupportQMCItemExtIEsEntryPresent = iota
	SliceSupportQMCItemExtIEsEntryPresentUnknown
)

type SliceSupportQMCItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SliceSupportQMCItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SliceSupportQMCItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SliceSupportQMC-Item-ExtIEsEntry", v)
}

// DecodeSliceSupportQMCItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSliceSupportQMCItemExtIEsEntry(data []byte) (SliceSupportQMCItemExtIEsEntry, int, error) {
	var v SliceSupportQMCItemExtIEsEntry
	err := ngapcodec.DecodeParameter("SliceSupportQMC-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SliceSupportQMCItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
