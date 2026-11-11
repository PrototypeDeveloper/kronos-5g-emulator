// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type DRBsSubjectToStatusTransferItemExtIEsEntryPresent uint32

const (
	DRBsSubjectToStatusTransferItemExtIEsEntryPresentNothing DRBsSubjectToStatusTransferItemExtIEsEntryPresent = iota
	DRBsSubjectToStatusTransferItemExtIEsEntryPresentOldAssociatedQosFlowListULendmarkerexpected
	DRBsSubjectToStatusTransferItemExtIEsEntryPresentUnknown
)

type DRBsSubjectToStatusTransferItemExtIEsEntry struct {
	ID                                          ProtocolExtensionID
	Criticality                                 Criticality
	Present                                     DRBsSubjectToStatusTransferItemExtIEsEntryPresent
	OldAssociatedQosFlowListULendmarkerexpected AssociatedQosFlowList `ngap:"id:159,criticality:ignore,presence:optional"`
	UnknownValue                                []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBsSubjectToStatusTransferItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DRBsSubjectToStatusTransferItem-ExtIEsEntry", v)
}

// DecodeDRBsSubjectToStatusTransferItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDRBsSubjectToStatusTransferItemExtIEsEntry(data []byte) (DRBsSubjectToStatusTransferItemExtIEsEntry, int, error) {
	var v DRBsSubjectToStatusTransferItemExtIEsEntry
	err := ngapcodec.DecodeParameter("DRBsSubjectToStatusTransferItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBsSubjectToStatusTransferItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
