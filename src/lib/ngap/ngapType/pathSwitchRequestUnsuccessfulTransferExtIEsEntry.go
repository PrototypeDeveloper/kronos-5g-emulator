// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PathSwitchRequestUnsuccessfulTransferExtIEsEntryPresent uint32

const (
	PathSwitchRequestUnsuccessfulTransferExtIEsEntryPresentNothing PathSwitchRequestUnsuccessfulTransferExtIEsEntryPresent = iota
	PathSwitchRequestUnsuccessfulTransferExtIEsEntryPresentUnknown
)

type PathSwitchRequestUnsuccessfulTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PathSwitchRequestUnsuccessfulTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PathSwitchRequestUnsuccessfulTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PathSwitchRequestUnsuccessfulTransfer-ExtIEsEntry", v)
}

// DecodePathSwitchRequestUnsuccessfulTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePathSwitchRequestUnsuccessfulTransferExtIEsEntry(data []byte) (PathSwitchRequestUnsuccessfulTransferExtIEsEntry, int, error) {
	var v PathSwitchRequestUnsuccessfulTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("PathSwitchRequestUnsuccessfulTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PathSwitchRequestUnsuccessfulTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
