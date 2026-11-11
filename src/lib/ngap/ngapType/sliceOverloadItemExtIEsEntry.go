// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SliceOverloadItemExtIEsEntryPresent uint32

const (
	SliceOverloadItemExtIEsEntryPresentNothing SliceOverloadItemExtIEsEntryPresent = iota
	SliceOverloadItemExtIEsEntryPresentUnknown
)

type SliceOverloadItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SliceOverloadItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SliceOverloadItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SliceOverloadItem-ExtIEsEntry", v)
}

// DecodeSliceOverloadItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSliceOverloadItemExtIEsEntry(data []byte) (SliceOverloadItemExtIEsEntry, int, error) {
	var v SliceOverloadItemExtIEsEntry
	err := ngapcodec.DecodeParameter("SliceOverloadItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SliceOverloadItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
