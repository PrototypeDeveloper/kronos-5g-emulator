// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSSessionFailedtoSetupItemExtIEsEntryPresent uint32

const (
	MBSSessionFailedtoSetupItemExtIEsEntryPresentNothing MBSSessionFailedtoSetupItemExtIEsEntryPresent = iota
	MBSSessionFailedtoSetupItemExtIEsEntryPresentUnknown
)

type MBSSessionFailedtoSetupItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSSessionFailedtoSetupItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionFailedtoSetupItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionFailedtoSetupItem-ExtIEsEntry", v)
}

// DecodeMBSSessionFailedtoSetupItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionFailedtoSetupItemExtIEsEntry(data []byte) (MBSSessionFailedtoSetupItemExtIEsEntry, int, error) {
	var v MBSSessionFailedtoSetupItemExtIEsEntry
	err := ngapcodec.DecodeParameter("MBSSessionFailedtoSetupItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionFailedtoSetupItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
