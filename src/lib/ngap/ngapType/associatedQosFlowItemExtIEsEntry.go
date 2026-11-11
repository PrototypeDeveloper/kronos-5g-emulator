// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AssociatedQosFlowItemExtIEsEntryPresent uint32

const (
	AssociatedQosFlowItemExtIEsEntryPresentNothing AssociatedQosFlowItemExtIEsEntryPresent = iota
	AssociatedQosFlowItemExtIEsEntryPresentCurrentQoSParaSetIndex
	AssociatedQosFlowItemExtIEsEntryPresentUnknown
)

type AssociatedQosFlowItemExtIEsEntry struct {
	ID                     ProtocolExtensionID
	Criticality            Criticality
	Present                AssociatedQosFlowItemExtIEsEntryPresent
	CurrentQoSParaSetIndex AlternativeQoSParaSetIndex `ngap:"id:221,criticality:ignore,presence:optional"`
	UnknownValue           []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AssociatedQosFlowItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AssociatedQosFlowItem-ExtIEsEntry", v)
}

// DecodeAssociatedQosFlowItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAssociatedQosFlowItemExtIEsEntry(data []byte) (AssociatedQosFlowItemExtIEsEntry, int, error) {
	var v AssociatedQosFlowItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AssociatedQosFlowItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AssociatedQosFlowItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
