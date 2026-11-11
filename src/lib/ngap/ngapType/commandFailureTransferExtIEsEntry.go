// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CommandFailureTransferExtIEsEntryPresent uint32

const (
	CommandFailureTransferExtIEsEntryPresentNothing CommandFailureTransferExtIEsEntryPresent = iota
	CommandFailureTransferExtIEsEntryPresentUnknown
)

type CommandFailureTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CommandFailureTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CommandFailureTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CommandFailureTransfer-ExtIEsEntry", v)
}

// DecodeCommandFailureTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCommandFailureTransferExtIEsEntry(data []byte) (CommandFailureTransferExtIEsEntry, int, error) {
	var v CommandFailureTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("CommandFailureTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CommandFailureTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
