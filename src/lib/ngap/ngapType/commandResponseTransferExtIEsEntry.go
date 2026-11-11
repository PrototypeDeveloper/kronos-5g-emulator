// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CommandResponseTransferExtIEsEntryPresent uint32

const (
	CommandResponseTransferExtIEsEntryPresentNothing CommandResponseTransferExtIEsEntryPresent = iota
	CommandResponseTransferExtIEsEntryPresentUnknown
)

type CommandResponseTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CommandResponseTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CommandResponseTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CommandResponseTransfer-ExtIEsEntry", v)
}

// DecodeCommandResponseTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCommandResponseTransferExtIEsEntry(data []byte) (CommandResponseTransferExtIEsEntry, int, error) {
	var v CommandResponseTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("CommandResponseTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CommandResponseTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
