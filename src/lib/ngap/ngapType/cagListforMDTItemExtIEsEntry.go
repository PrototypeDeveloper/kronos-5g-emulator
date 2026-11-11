// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CAGListforMDTItemExtIEsEntryPresent uint32

const (
	CAGListforMDTItemExtIEsEntryPresentNothing CAGListforMDTItemExtIEsEntryPresent = iota
	CAGListforMDTItemExtIEsEntryPresentUnknown
)

type CAGListforMDTItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CAGListforMDTItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CAGListforMDTItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CAGListforMDTItem-ExtIEsEntry", v)
}

// DecodeCAGListforMDTItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCAGListforMDTItemExtIEsEntry(data []byte) (CAGListforMDTItemExtIEsEntry, int, error) {
	var v CAGListforMDTItemExtIEsEntry
	err := ngapcodec.DecodeParameter("CAGListforMDTItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CAGListforMDTItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
