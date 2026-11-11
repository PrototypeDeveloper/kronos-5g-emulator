// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSSessionTNLInfo5GCExtIEsEntryPresent uint32

const (
	MBSSessionTNLInfo5GCExtIEsEntryPresentNothing MBSSessionTNLInfo5GCExtIEsEntryPresent = iota
	MBSSessionTNLInfo5GCExtIEsEntryPresentUnknown
)

type MBSSessionTNLInfo5GCExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      MBSSessionTNLInfo5GCExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionTNLInfo5GCExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-SessionTNLInfo5GC-ExtIEsEntry", v)
}

// DecodeMBSSessionTNLInfo5GCExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionTNLInfo5GCExtIEsEntry(data []byte) (MBSSessionTNLInfo5GCExtIEsEntry, int, error) {
	var v MBSSessionTNLInfo5GCExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-SessionTNLInfo5GC-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionTNLInfo5GCExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
