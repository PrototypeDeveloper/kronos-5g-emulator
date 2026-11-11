// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type HandoverRequestAcknowledgeTransferExtIEsEntryPresent uint32

const (
	HandoverRequestAcknowledgeTransferExtIEsEntryPresentNothing HandoverRequestAcknowledgeTransferExtIEsEntryPresent = iota
	HandoverRequestAcknowledgeTransferExtIEsEntryPresentAdditionalDLUPTNLInformationForHOList
	HandoverRequestAcknowledgeTransferExtIEsEntryPresentULForwardingUPTNLInformation
	HandoverRequestAcknowledgeTransferExtIEsEntryPresentAdditionalULForwardingUPTNLInformation
	HandoverRequestAcknowledgeTransferExtIEsEntryPresentDataForwardingResponseERABList
	HandoverRequestAcknowledgeTransferExtIEsEntryPresentRedundantDLNGUUPTNLInformation
	HandoverRequestAcknowledgeTransferExtIEsEntryPresentUsedRSNInformation
	HandoverRequestAcknowledgeTransferExtIEsEntryPresentGlobalRANNodeID
	HandoverRequestAcknowledgeTransferExtIEsEntryPresentMBSSupportIndicator
	HandoverRequestAcknowledgeTransferExtIEsEntryPresentPDUSetbasedHandlingIndicator
	HandoverRequestAcknowledgeTransferExtIEsEntryPresentECNMarkingorCongestionInformationReportingStatus
	HandoverRequestAcknowledgeTransferExtIEsEntryPresentUnknown
)

type HandoverRequestAcknowledgeTransferExtIEsEntry struct {
	ID                                               ProtocolExtensionID
	Criticality                                      Criticality
	Present                                          HandoverRequestAcknowledgeTransferExtIEsEntryPresent
	AdditionalDLUPTNLInformationForHOList            AdditionalDLUPTNLInformationForHOList            `ngap:"id:153,criticality:ignore,presence:optional"`
	ULForwardingUPTNLInformation                     UPTransportLayerInformation                      `ngap:"id:164,criticality:reject,presence:optional"`
	AdditionalULForwardingUPTNLInformation           UPTransportLayerInformationList                  `ngap:"id:172,criticality:reject,presence:optional"`
	DataForwardingResponseERABList                   DataForwardingResponseERABList                   `ngap:"id:249,criticality:ignore,presence:optional"`
	RedundantDLNGUUPTNLInformation                   UPTransportLayerInformation                      `ngap:"id:192,criticality:ignore,presence:optional"`
	UsedRSNInformation                               RedundantPDUSessionInformation                   `ngap:"id:198,criticality:ignore,presence:optional"`
	GlobalRANNodeID                                  GlobalRANNodeID                                  `ngap:"id:27,criticality:ignore,presence:optional"`
	MBSSupportIndicator                              MBSSupportIndicator                              `ngap:"id:309,criticality:ignore,presence:optional"`
	PDUSetbasedHandlingIndicator                     PDUSetbasedHandlingIndicator                     `ngap:"id:423,criticality:ignore,presence:optional"`
	ECNMarkingorCongestionInformationReportingStatus ECNMarkingorCongestionInformationReportingStatus `ngap:"id:426,criticality:ignore,presence:optional"`
	UnknownValue                                     []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HandoverRequestAcknowledgeTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("HandoverRequestAcknowledgeTransfer-ExtIEsEntry", v)
}

// DecodeHandoverRequestAcknowledgeTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeHandoverRequestAcknowledgeTransferExtIEsEntry(data []byte) (HandoverRequestAcknowledgeTransferExtIEsEntry, int, error) {
	var v HandoverRequestAcknowledgeTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("HandoverRequestAcknowledgeTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return HandoverRequestAcknowledgeTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
