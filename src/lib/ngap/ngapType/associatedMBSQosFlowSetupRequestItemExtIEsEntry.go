// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AssociatedMBSQosFlowSetupRequestItemExtIEsEntryPresent uint32

const (
	AssociatedMBSQosFlowSetupRequestItemExtIEsEntryPresentNothing AssociatedMBSQosFlowSetupRequestItemExtIEsEntryPresent = iota
	AssociatedMBSQosFlowSetupRequestItemExtIEsEntryPresentUnknown
)

type AssociatedMBSQosFlowSetupRequestItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AssociatedMBSQosFlowSetupRequestItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AssociatedMBSQosFlowSetupRequestItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AssociatedMBSQosFlowSetupRequestItem-ExtIEsEntry", v)
}

// DecodeAssociatedMBSQosFlowSetupRequestItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAssociatedMBSQosFlowSetupRequestItemExtIEsEntry(data []byte) (AssociatedMBSQosFlowSetupRequestItemExtIEsEntry, int, error) {
	var v AssociatedMBSQosFlowSetupRequestItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AssociatedMBSQosFlowSetupRequestItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AssociatedMBSQosFlowSetupRequestItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
