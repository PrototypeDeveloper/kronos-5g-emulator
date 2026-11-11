// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceModifyResponseTransferExtIEsEntryPresent uint32

const (
	PDUSessionResourceModifyResponseTransferExtIEsEntryPresentNothing PDUSessionResourceModifyResponseTransferExtIEsEntryPresent = iota
	PDUSessionResourceModifyResponseTransferExtIEsEntryPresentAdditionalNGUUPTNLInformation
	PDUSessionResourceModifyResponseTransferExtIEsEntryPresentRedundantDLNGUUPTNLInformation
	PDUSessionResourceModifyResponseTransferExtIEsEntryPresentRedundantULNGUUPTNLInformation
	PDUSessionResourceModifyResponseTransferExtIEsEntryPresentAdditionalRedundantDLQosFlowPerTNLInformation
	PDUSessionResourceModifyResponseTransferExtIEsEntryPresentAdditionalRedundantNGUUPTNLInformation
	PDUSessionResourceModifyResponseTransferExtIEsEntryPresentSecondaryRATUsageInformation
	PDUSessionResourceModifyResponseTransferExtIEsEntryPresentMBSSupportIndicator
	PDUSessionResourceModifyResponseTransferExtIEsEntryPresentMBSSessionSetuporModifyResponseList
	PDUSessionResourceModifyResponseTransferExtIEsEntryPresentMBSSessionFailedtoSetuporModifyList
	PDUSessionResourceModifyResponseTransferExtIEsEntryPresentECNMarkingorCongestionInformationReportingStatus
	PDUSessionResourceModifyResponseTransferExtIEsEntryPresentPDUSetbasedHandlingIndicator
	PDUSessionResourceModifyResponseTransferExtIEsEntryPresentUserPlaneFailureIndicationReport
	PDUSessionResourceModifyResponseTransferExtIEsEntryPresentUnknown
)

type PDUSessionResourceModifyResponseTransferExtIEsEntry struct {
	ID                                               ProtocolExtensionID
	Criticality                                      Criticality
	Present                                          PDUSessionResourceModifyResponseTransferExtIEsEntryPresent
	AdditionalNGUUPTNLInformation                    UPTransportLayerInformationPairList              `ngap:"id:154,criticality:ignore,presence:optional"`
	RedundantDLNGUUPTNLInformation                   UPTransportLayerInformation                      `ngap:"id:192,criticality:ignore,presence:optional"`
	RedundantULNGUUPTNLInformation                   UPTransportLayerInformation                      `ngap:"id:195,criticality:ignore,presence:optional"`
	AdditionalRedundantDLQosFlowPerTNLInformation    QosFlowPerTNLInformationList                     `ngap:"id:184,criticality:ignore,presence:optional"`
	AdditionalRedundantNGUUPTNLInformation           UPTransportLayerInformationPairList              `ngap:"id:185,criticality:ignore,presence:optional"`
	SecondaryRATUsageInformation                     SecondaryRATUsageInformation                     `ngap:"id:144,criticality:ignore,presence:optional"`
	MBSSupportIndicator                              MBSSupportIndicator                              `ngap:"id:309,criticality:ignore,presence:optional"`
	MBSSessionSetuporModifyResponseList              MBSSessionSetupResponseList                      `ngap:"id:313,criticality:ignore,presence:optional"`
	MBSSessionFailedtoSetuporModifyList              MBSSessionFailedtoSetupList                      `ngap:"id:311,criticality:ignore,presence:optional"`
	ECNMarkingorCongestionInformationReportingStatus ECNMarkingorCongestionInformationReportingStatus `ngap:"id:426,criticality:ignore,presence:optional"`
	PDUSetbasedHandlingIndicator                     PDUSetbasedHandlingIndicator                     `ngap:"id:423,criticality:ignore,presence:optional"`
	UserPlaneFailureIndicationReport                 UserPlaneFailureIndicationReport                 `ngap:"id:436,criticality:ignore,presence:optional"`
	UnknownValue                                     []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyResponseTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyResponseTransfer-ExtIEsEntry", v)
}

// DecodePDUSessionResourceModifyResponseTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyResponseTransferExtIEsEntry(data []byte) (PDUSessionResourceModifyResponseTransferExtIEsEntry, int, error) {
	var v PDUSessionResourceModifyResponseTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyResponseTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyResponseTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
