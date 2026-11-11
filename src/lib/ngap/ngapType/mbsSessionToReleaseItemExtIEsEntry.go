// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSSessionToReleaseItemExtIEsEntryPresent uint32

const (
	MBSSessionToReleaseItemExtIEsEntryPresentNothing MBSSessionToReleaseItemExtIEsEntryPresent = iota
	MBSSessionToReleaseItemExtIEsEntryPresentUnknown
)

type MBSSessionToReleaseItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSSessionToReleaseItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionToReleaseItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionToReleaseItem-ExtIEsEntry", v)
}

// DecodeMBSSessionToReleaseItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionToReleaseItemExtIEsEntry(data []byte) (MBSSessionToReleaseItemExtIEsEntry, int, error) {
	var v MBSSessionToReleaseItemExtIEsEntry
	err := ngapcodec.DecodeParameter("MBSSessionToReleaseItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionToReleaseItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
