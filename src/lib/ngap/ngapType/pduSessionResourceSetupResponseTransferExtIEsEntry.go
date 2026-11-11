// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceSetupResponseTransferExtIEsEntryPresent uint32

const (
	PDUSessionResourceSetupResponseTransferExtIEsEntryPresentNothing PDUSessionResourceSetupResponseTransferExtIEsEntryPresent = iota
	PDUSessionResourceSetupResponseTransferExtIEsEntryPresentRedundantDLQosFlowPerTNLInformation
	PDUSessionResourceSetupResponseTransferExtIEsEntryPresentAdditionalRedundantDLQosFlowPerTNLInformation
	PDUSessionResourceSetupResponseTransferExtIEsEntryPresentUsedRSNInformation
	PDUSessionResourceSetupResponseTransferExtIEsEntryPresentGlobalRANNodeID
	PDUSessionResourceSetupResponseTransferExtIEsEntryPresentMBSSupportIndicator
	PDUSessionResourceSetupResponseTransferExtIEsEntryPresentMBSSessionSetupResponseList
	PDUSessionResourceSetupResponseTransferExtIEsEntryPresentMBSSessionFailedtoSetupList
	PDUSessionResourceSetupResponseTransferExtIEsEntryPresentQoSFlowTSCList
	PDUSessionResourceSetupResponseTransferExtIEsEntryPresentUplinkTLContainer
	PDUSessionResourceSetupResponseTransferExtIEsEntryPresentECNMarkingorCongestionInformationReportingStatus
	PDUSessionResourceSetupResponseTransferExtIEsEntryPresentPDUSetbasedHandlingIndicator
	PDUSessionResourceSetupResponseTransferExtIEsEntryPresentUnknown
)

type PDUSessionResourceSetupResponseTransferExtIEsEntry struct {
	ID                                               ProtocolExtensionID
	Criticality                                      Criticality
	Present                                          PDUSessionResourceSetupResponseTransferExtIEsEntryPresent
	RedundantDLQosFlowPerTNLInformation              QosFlowPerTNLInformation                         `ngap:"id:193,criticality:ignore,presence:optional"`
	AdditionalRedundantDLQosFlowPerTNLInformation    QosFlowPerTNLInformationList                     `ngap:"id:184,criticality:ignore,presence:optional"`
	UsedRSNInformation                               RedundantPDUSessionInformation                   `ngap:"id:198,criticality:ignore,presence:optional"`
	GlobalRANNodeID                                  GlobalRANNodeID                                  `ngap:"id:27,criticality:ignore,presence:optional"`
	MBSSupportIndicator                              MBSSupportIndicator                              `ngap:"id:309,criticality:ignore,presence:optional"`
	MBSSessionSetupResponseList                      MBSSessionSetupResponseList                      `ngap:"id:312,criticality:ignore,presence:optional"`
	MBSSessionFailedtoSetupList                      MBSSessionFailedtoSetupList                      `ngap:"id:310,criticality:ignore,presence:optional"`
	QoSFlowTSCList                                   QoSFlowTSCList                                   `ngap:"id:392,criticality:ignore,presence:optional"`
	UplinkTLContainer                                TLContainer                                      `ngap:"id:395,criticality:ignore,presence:optional"`
	ECNMarkingorCongestionInformationReportingStatus ECNMarkingorCongestionInformationReportingStatus `ngap:"id:426,criticality:ignore,presence:optional"`
	PDUSetbasedHandlingIndicator                     PDUSetbasedHandlingIndicator                     `ngap:"id:423,criticality:ignore,presence:optional"`
	UnknownValue                                     []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupResponseTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupResponseTransfer-ExtIEsEntry", v)
}

// DecodePDUSessionResourceSetupResponseTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupResponseTransferExtIEsEntry(data []byte) (PDUSessionResourceSetupResponseTransferExtIEsEntry, int, error) {
	var v PDUSessionResourceSetupResponseTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupResponseTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupResponseTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
