// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSSessionSetupOrModFailureTransferExtIEsEntryPresent uint32

const (
	MBSSessionSetupOrModFailureTransferExtIEsEntryPresentNothing MBSSessionSetupOrModFailureTransferExtIEsEntryPresent = iota
	MBSSessionSetupOrModFailureTransferExtIEsEntryPresentUnknown
)

type MBSSessionSetupOrModFailureTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSSessionSetupOrModFailureTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionSetupOrModFailureTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionSetupOrModFailureTransfer-ExtIEsEntry", v)
}

// DecodeMBSSessionSetupOrModFailureTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionSetupOrModFailureTransferExtIEsEntry(data []byte) (MBSSessionSetupOrModFailureTransferExtIEsEntry, int, error) {
	var v MBSSessionSetupOrModFailureTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("MBSSessionSetupOrModFailureTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionSetupOrModFailureTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
