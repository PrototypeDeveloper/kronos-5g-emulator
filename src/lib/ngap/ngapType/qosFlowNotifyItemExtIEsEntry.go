// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QosFlowNotifyItemExtIEsEntryPresent uint32

const (
	QosFlowNotifyItemExtIEsEntryPresentNothing QosFlowNotifyItemExtIEsEntryPresent = iota
	QosFlowNotifyItemExtIEsEntryPresentCurrentQoSParaSetIndex
	QosFlowNotifyItemExtIEsEntryPresentTSCTrafficCharacteristicsFeedback
	QosFlowNotifyItemExtIEsEntryPresentUnknown
)

type QosFlowNotifyItemExtIEsEntry struct {
	ID                                ProtocolExtensionID
	Criticality                       Criticality
	Present                           QosFlowNotifyItemExtIEsEntryPresent
	CurrentQoSParaSetIndex            AlternativeQoSParaSetNotifyIndex  `ngap:"id:221,criticality:ignore,presence:optional"`
	TSCTrafficCharacteristicsFeedback TSCTrafficCharacteristicsFeedback `ngap:"id:393,criticality:ignore,presence:optional"`
	UnknownValue                      []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowNotifyItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowNotifyItem-ExtIEsEntry", v)
}

// DecodeQosFlowNotifyItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowNotifyItemExtIEsEntry(data []byte) (QosFlowNotifyItemExtIEsEntry, int, error) {
	var v QosFlowNotifyItemExtIEsEntry
	err := ngapcodec.DecodeParameter("QosFlowNotifyItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowNotifyItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
