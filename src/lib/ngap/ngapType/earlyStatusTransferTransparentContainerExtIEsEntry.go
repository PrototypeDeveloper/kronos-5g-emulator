// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EarlyStatusTransferTransparentContainerExtIEsEntryPresent uint32

const (
	EarlyStatusTransferTransparentContainerExtIEsEntryPresentNothing EarlyStatusTransferTransparentContainerExtIEsEntryPresent = iota
	EarlyStatusTransferTransparentContainerExtIEsEntryPresentUnknown
)

type EarlyStatusTransferTransparentContainerExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EarlyStatusTransferTransparentContainerExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EarlyStatusTransferTransparentContainerExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EarlyStatusTransfer-TransparentContainer-ExtIEsEntry", v)
}

// DecodeEarlyStatusTransferTransparentContainerExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEarlyStatusTransferTransparentContainerExtIEsEntry(data []byte) (EarlyStatusTransferTransparentContainerExtIEsEntry, int, error) {
	var v EarlyStatusTransferTransparentContainerExtIEsEntry
	err := ngapcodec.DecodeParameter("EarlyStatusTransfer-TransparentContainer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EarlyStatusTransferTransparentContainerExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
