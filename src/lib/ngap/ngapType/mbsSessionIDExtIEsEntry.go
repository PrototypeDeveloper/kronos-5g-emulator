// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSSessionIDExtIEsEntryPresent uint32

const (
	MBSSessionIDExtIEsEntryPresentNothing MBSSessionIDExtIEsEntryPresent = iota
	MBSSessionIDExtIEsEntryPresentUnknown
)

type MBSSessionIDExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSSessionIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-SessionID-ExtIEsEntry", v)
}

// DecodeMBSSessionIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionIDExtIEsEntry(data []byte) (MBSSessionIDExtIEsEntry, int, error) {
	var v MBSSessionIDExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-SessionID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
