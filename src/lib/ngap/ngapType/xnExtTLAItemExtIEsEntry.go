// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type XnExtTLAItemExtIEsEntryPresent uint32

const (
	XnExtTLAItemExtIEsEntryPresentNothing XnExtTLAItemExtIEsEntryPresent = iota
	XnExtTLAItemExtIEsEntryPresentSCTPTLAs
	XnExtTLAItemExtIEsEntryPresentUnknown
)

type XnExtTLAItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      XnExtTLAItemExtIEsEntryPresent
	SCTPTLAs     SCTPTLAs `ngap:"id:173,criticality:ignore,presence:optional"`
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v XnExtTLAItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("XnExtTLA-Item-ExtIEsEntry", v)
}

// DecodeXnExtTLAItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeXnExtTLAItemExtIEsEntry(data []byte) (XnExtTLAItemExtIEsEntry, int, error) {
	var v XnExtTLAItemExtIEsEntry
	err := ngapcodec.DecodeParameter("XnExtTLA-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return XnExtTLAItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
