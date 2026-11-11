// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RANStatusTransferTransparentContainerExtIEsEntryPresent uint32

const (
	RANStatusTransferTransparentContainerExtIEsEntryPresentNothing RANStatusTransferTransparentContainerExtIEsEntryPresent = iota
	RANStatusTransferTransparentContainerExtIEsEntryPresentUnknown
)

type RANStatusTransferTransparentContainerExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      RANStatusTransferTransparentContainerExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANStatusTransferTransparentContainerExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RANStatusTransfer-TransparentContainer-ExtIEsEntry", v)
}

// DecodeRANStatusTransferTransparentContainerExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRANStatusTransferTransparentContainerExtIEsEntry(data []byte) (RANStatusTransferTransparentContainerExtIEsEntry, int, error) {
	var v RANStatusTransferTransparentContainerExtIEsEntry
	err := ngapcodec.DecodeParameter("RANStatusTransfer-TransparentContainer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANStatusTransferTransparentContainerExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
