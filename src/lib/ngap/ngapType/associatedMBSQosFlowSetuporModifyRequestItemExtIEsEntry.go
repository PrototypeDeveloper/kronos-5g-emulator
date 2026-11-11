// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AssociatedMBSQosFlowSetuporModifyRequestItemExtIEsEntryPresent uint32

const (
	AssociatedMBSQosFlowSetuporModifyRequestItemExtIEsEntryPresentNothing AssociatedMBSQosFlowSetuporModifyRequestItemExtIEsEntryPresent = iota
	AssociatedMBSQosFlowSetuporModifyRequestItemExtIEsEntryPresentUnknown
)

type AssociatedMBSQosFlowSetuporModifyRequestItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AssociatedMBSQosFlowSetuporModifyRequestItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AssociatedMBSQosFlowSetuporModifyRequestItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AssociatedMBSQosFlowSetuporModifyRequestItem-ExtIEsEntry", v)
}

// DecodeAssociatedMBSQosFlowSetuporModifyRequestItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAssociatedMBSQosFlowSetuporModifyRequestItemExtIEsEntry(data []byte) (AssociatedMBSQosFlowSetuporModifyRequestItemExtIEsEntry, int, error) {
	var v AssociatedMBSQosFlowSetuporModifyRequestItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AssociatedMBSQosFlowSetuporModifyRequestItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AssociatedMBSQosFlowSetuporModifyRequestItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
