// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceModifyIndicationTransferExtIEsEntryPresent uint32

const (
	PDUSessionResourceModifyIndicationTransferExtIEsEntryPresentNothing PDUSessionResourceModifyIndicationTransferExtIEsEntryPresent = iota
	PDUSessionResourceModifyIndicationTransferExtIEsEntryPresentSecondaryRATUsageInformation
	PDUSessionResourceModifyIndicationTransferExtIEsEntryPresentSecurityResult
	PDUSessionResourceModifyIndicationTransferExtIEsEntryPresentRedundantDLQosFlowPerTNLInformation
	PDUSessionResourceModifyIndicationTransferExtIEsEntryPresentAdditionalRedundantDLQosFlowPerTNLInformation
	PDUSessionResourceModifyIndicationTransferExtIEsEntryPresentGlobalRANNodeID
	PDUSessionResourceModifyIndicationTransferExtIEsEntryPresentECNMarkingorCongestionInformationReportingStatus
	PDUSessionResourceModifyIndicationTransferExtIEsEntryPresentUnknown
)

type PDUSessionResourceModifyIndicationTransferExtIEsEntry struct {
	ID                                               ProtocolExtensionID
	Criticality                                      Criticality
	Present                                          PDUSessionResourceModifyIndicationTransferExtIEsEntryPresent
	SecondaryRATUsageInformation                     SecondaryRATUsageInformation                     `ngap:"id:144,criticality:ignore,presence:optional"`
	SecurityResult                                   SecurityResult                                   `ngap:"id:156,criticality:ignore,presence:optional"`
	RedundantDLQosFlowPerTNLInformation              QosFlowPerTNLInformation                         `ngap:"id:193,criticality:ignore,presence:optional"`
	AdditionalRedundantDLQosFlowPerTNLInformation    QosFlowPerTNLInformationList                     `ngap:"id:184,criticality:ignore,presence:optional"`
	GlobalRANNodeID                                  GlobalRANNodeID                                  `ngap:"id:27,criticality:ignore,presence:optional"`
	ECNMarkingorCongestionInformationReportingStatus ECNMarkingorCongestionInformationReportingStatus `ngap:"id:426,criticality:ignore,presence:optional"`
	UnknownValue                                     []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyIndicationTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyIndicationTransfer-ExtIEsEntry", v)
}

// DecodePDUSessionResourceModifyIndicationTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyIndicationTransferExtIEsEntry(data []byte) (PDUSessionResourceModifyIndicationTransferExtIEsEntry, int, error) {
	var v PDUSessionResourceModifyIndicationTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyIndicationTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyIndicationTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
