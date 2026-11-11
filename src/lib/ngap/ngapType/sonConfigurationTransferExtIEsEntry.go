// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SONConfigurationTransferExtIEsEntryPresent uint32

const (
	SONConfigurationTransferExtIEsEntryPresentNothing SONConfigurationTransferExtIEsEntryPresent = iota
	SONConfigurationTransferExtIEsEntryPresentRequestedTNLInfo
	SONConfigurationTransferExtIEsEntryPresentUnknown
)

type SONConfigurationTransferExtIEsEntry struct {
	ID               ProtocolExtensionID
	Criticality      Criticality
	Present          SONConfigurationTransferExtIEsEntryPresent
	RequestedTNLInfo RequestedTNLInfo `ngap:"id:445,criticality:ignore,presence:optional"`
	UnknownValue     []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SONConfigurationTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SONConfigurationTransfer-ExtIEsEntry", v)
}

// DecodeSONConfigurationTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSONConfigurationTransferExtIEsEntry(data []byte) (SONConfigurationTransferExtIEsEntry, int, error) {
	var v SONConfigurationTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("SONConfigurationTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SONConfigurationTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
