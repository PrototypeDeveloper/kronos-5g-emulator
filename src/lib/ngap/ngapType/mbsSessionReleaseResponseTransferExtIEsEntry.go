// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSSessionReleaseResponseTransferExtIEsEntryPresent uint32

const (
	MBSSessionReleaseResponseTransferExtIEsEntryPresentNothing MBSSessionReleaseResponseTransferExtIEsEntryPresent = iota
	MBSSessionReleaseResponseTransferExtIEsEntryPresentUnknown
)

type MBSSessionReleaseResponseTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSSessionReleaseResponseTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionReleaseResponseTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionReleaseResponseTransfer-ExtIEsEntry", v)
}

// DecodeMBSSessionReleaseResponseTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionReleaseResponseTransferExtIEsEntry(data []byte) (MBSSessionReleaseResponseTransferExtIEsEntry, int, error) {
	var v MBSSessionReleaseResponseTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("MBSSessionReleaseResponseTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionReleaseResponseTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
