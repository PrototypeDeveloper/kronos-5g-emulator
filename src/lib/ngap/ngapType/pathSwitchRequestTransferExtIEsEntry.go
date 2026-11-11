// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PathSwitchRequestTransferExtIEsEntryPresent uint32

const (
	PathSwitchRequestTransferExtIEsEntryPresentNothing PathSwitchRequestTransferExtIEsEntryPresent = iota
	PathSwitchRequestTransferExtIEsEntryPresentAdditionalDLQosFlowPerTNLInformation
	PathSwitchRequestTransferExtIEsEntryPresentRedundantDLNGUUPTNLInformation
	PathSwitchRequestTransferExtIEsEntryPresentRedundantDLNGUTNLInformationReused
	PathSwitchRequestTransferExtIEsEntryPresentAdditionalRedundantDLQosFlowPerTNLInformation
	PathSwitchRequestTransferExtIEsEntryPresentUsedRSNInformation
	PathSwitchRequestTransferExtIEsEntryPresentGlobalRANNodeID
	PathSwitchRequestTransferExtIEsEntryPresentMBSSupportIndicator
	PathSwitchRequestTransferExtIEsEntryPresentPDUSetbasedHandlingIndicator
	PathSwitchRequestTransferExtIEsEntryPresentECNMarkingorCongestionInformationReportingStatus
	PathSwitchRequestTransferExtIEsEntryPresentUnknown
)

type PathSwitchRequestTransferExtIEsEntry struct {
	ID                                               ProtocolExtensionID
	Criticality                                      Criticality
	Present                                          PathSwitchRequestTransferExtIEsEntryPresent
	AdditionalDLQosFlowPerTNLInformation             QosFlowPerTNLInformationList                     `ngap:"id:155,criticality:ignore,presence:optional"`
	RedundantDLNGUUPTNLInformation                   UPTransportLayerInformation                      `ngap:"id:192,criticality:ignore,presence:optional"`
	RedundantDLNGUTNLInformationReused               DLNGUTNLInformationReused                        `ngap:"id:191,criticality:ignore,presence:optional"`
	AdditionalRedundantDLQosFlowPerTNLInformation    QosFlowPerTNLInformationList                     `ngap:"id:184,criticality:ignore,presence:optional"`
	UsedRSNInformation                               RedundantPDUSessionInformation                   `ngap:"id:198,criticality:ignore,presence:optional"`
	GlobalRANNodeID                                  GlobalRANNodeID                                  `ngap:"id:27,criticality:ignore,presence:optional"`
	MBSSupportIndicator                              MBSSupportIndicator                              `ngap:"id:309,criticality:ignore,presence:optional"`
	PDUSetbasedHandlingIndicator                     PDUSetbasedHandlingIndicator                     `ngap:"id:423,criticality:ignore,presence:optional"`
	ECNMarkingorCongestionInformationReportingStatus ECNMarkingorCongestionInformationReportingStatus `ngap:"id:426,criticality:ignore,presence:optional"`
	UnknownValue                                     []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PathSwitchRequestTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PathSwitchRequestTransfer-ExtIEsEntry", v)
}

// DecodePathSwitchRequestTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePathSwitchRequestTransferExtIEsEntry(data []byte) (PathSwitchRequestTransferExtIEsEntry, int, error) {
	var v PathSwitchRequestTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("PathSwitchRequestTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PathSwitchRequestTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
