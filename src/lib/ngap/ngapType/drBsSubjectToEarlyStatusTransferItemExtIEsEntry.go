// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type DRBsSubjectToEarlyStatusTransferItemExtIEsEntryPresent uint32

const (
	DRBsSubjectToEarlyStatusTransferItemExtIEsEntryPresentNothing DRBsSubjectToEarlyStatusTransferItemExtIEsEntryPresent = iota
	DRBsSubjectToEarlyStatusTransferItemExtIEsEntryPresentUnknown
)

type DRBsSubjectToEarlyStatusTransferItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      DRBsSubjectToEarlyStatusTransferItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBsSubjectToEarlyStatusTransferItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DRBsSubjectToEarlyStatusTransfer-Item-ExtIEsEntry", v)
}

// DecodeDRBsSubjectToEarlyStatusTransferItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDRBsSubjectToEarlyStatusTransferItemExtIEsEntry(data []byte) (DRBsSubjectToEarlyStatusTransferItemExtIEsEntry, int, error) {
	var v DRBsSubjectToEarlyStatusTransferItemExtIEsEntry
	err := ngapcodec.DecodeParameter("DRBsSubjectToEarlyStatusTransfer-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBsSubjectToEarlyStatusTransferItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
