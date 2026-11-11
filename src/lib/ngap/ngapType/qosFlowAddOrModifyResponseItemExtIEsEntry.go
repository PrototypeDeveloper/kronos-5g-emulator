// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QosFlowAddOrModifyResponseItemExtIEsEntryPresent uint32

const (
	QosFlowAddOrModifyResponseItemExtIEsEntryPresentNothing QosFlowAddOrModifyResponseItemExtIEsEntryPresent = iota
	QosFlowAddOrModifyResponseItemExtIEsEntryPresentCurrentQoSParaSetIndex
	QosFlowAddOrModifyResponseItemExtIEsEntryPresentTSCTrafficCharacteristicsFeedback
	QosFlowAddOrModifyResponseItemExtIEsEntryPresentUplinkTLContainer
	QosFlowAddOrModifyResponseItemExtIEsEntryPresentDLNGUUPTNLInformation
	QosFlowAddOrModifyResponseItemExtIEsEntryPresentANPacketDelayBudgetUL
	QosFlowAddOrModifyResponseItemExtIEsEntryPresentUnknown
)

type QosFlowAddOrModifyResponseItemExtIEsEntry struct {
	ID                                ProtocolExtensionID
	Criticality                       Criticality
	Present                           QosFlowAddOrModifyResponseItemExtIEsEntryPresent
	CurrentQoSParaSetIndex            AlternativeQoSParaSetIndex        `ngap:"id:221,criticality:ignore,presence:optional"`
	TSCTrafficCharacteristicsFeedback TSCTrafficCharacteristicsFeedback `ngap:"id:393,criticality:ignore,presence:optional"`
	UplinkTLContainer                 TLContainer                       `ngap:"id:395,criticality:ignore,presence:optional"`
	DLNGUUPTNLInformation             UPTransportLayerInformation       `ngap:"id:128,criticality:ignore,presence:optional"`
	ANPacketDelayBudgetUL             ExtendedPacketDelayBudget         `ngap:"id:396,criticality:ignore,presence:optional"`
	UnknownValue                      []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowAddOrModifyResponseItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowAddOrModifyResponseItem-ExtIEsEntry", v)
}

// DecodeQosFlowAddOrModifyResponseItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowAddOrModifyResponseItemExtIEsEntry(data []byte) (QosFlowAddOrModifyResponseItemExtIEsEntry, int, error) {
	var v QosFlowAddOrModifyResponseItemExtIEsEntry
	err := ngapcodec.DecodeParameter("QosFlowAddOrModifyResponseItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowAddOrModifyResponseItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
