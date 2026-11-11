// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PathSwitchRequestSetupFailedTransferExtIEsEntryPresent uint32

const (
	PathSwitchRequestSetupFailedTransferExtIEsEntryPresentNothing PathSwitchRequestSetupFailedTransferExtIEsEntryPresent = iota
	PathSwitchRequestSetupFailedTransferExtIEsEntryPresentUnknown
)

type PathSwitchRequestSetupFailedTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PathSwitchRequestSetupFailedTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PathSwitchRequestSetupFailedTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PathSwitchRequestSetupFailedTransfer-ExtIEsEntry", v)
}

// DecodePathSwitchRequestSetupFailedTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePathSwitchRequestSetupFailedTransferExtIEsEntry(data []byte) (PathSwitchRequestSetupFailedTransferExtIEsEntry, int, error) {
	var v PathSwitchRequestSetupFailedTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("PathSwitchRequestSetupFailedTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PathSwitchRequestSetupFailedTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
