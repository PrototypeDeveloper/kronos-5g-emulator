// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSSessionTNLInfo5GCItemExtIEsEntryPresent uint32

const (
	MBSSessionTNLInfo5GCItemExtIEsEntryPresentNothing MBSSessionTNLInfo5GCItemExtIEsEntryPresent = iota
	MBSSessionTNLInfo5GCItemExtIEsEntryPresentUnknown
)

type MBSSessionTNLInfo5GCItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSSessionTNLInfo5GCItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionTNLInfo5GCItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-SessionTNLInfo5GCItem-ExtIEsEntry", v)
}

// DecodeMBSSessionTNLInfo5GCItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionTNLInfo5GCItemExtIEsEntry(data []byte) (MBSSessionTNLInfo5GCItemExtIEsEntry, int, error) {
	var v MBSSessionTNLInfo5GCItemExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-SessionTNLInfo5GCItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionTNLInfo5GCItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
