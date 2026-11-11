// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PathSwitchRequestAcknowledgeTransferExtIEsEntryPresent uint32

const (
	PathSwitchRequestAcknowledgeTransferExtIEsEntryPresentNothing PathSwitchRequestAcknowledgeTransferExtIEsEntryPresent = iota
	PathSwitchRequestAcknowledgeTransferExtIEsEntryPresentAdditionalNGUUPTNLInformation
	PathSwitchRequestAcknowledgeTransferExtIEsEntryPresentRedundantULNGUUPTNLInformation
	PathSwitchRequestAcknowledgeTransferExtIEsEntryPresentAdditionalRedundantNGUUPTNLInformation
	PathSwitchRequestAcknowledgeTransferExtIEsEntryPresentQosFlowParametersList
	PathSwitchRequestAcknowledgeTransferExtIEsEntryPresentUnknown
)

type PathSwitchRequestAcknowledgeTransferExtIEsEntry struct {
	ID                                     ProtocolExtensionID
	Criticality                            Criticality
	Present                                PathSwitchRequestAcknowledgeTransferExtIEsEntryPresent
	AdditionalNGUUPTNLInformation          UPTransportLayerInformationPairList `ngap:"id:154,criticality:ignore,presence:optional"`
	RedundantULNGUUPTNLInformation         UPTransportLayerInformation         `ngap:"id:195,criticality:ignore,presence:optional"`
	AdditionalRedundantNGUUPTNLInformation UPTransportLayerInformationPairList `ngap:"id:185,criticality:ignore,presence:optional"`
	QosFlowParametersList                  QosFlowParametersList               `ngap:"id:277,criticality:ignore,presence:optional"`
	UnknownValue                           []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PathSwitchRequestAcknowledgeTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PathSwitchRequestAcknowledgeTransfer-ExtIEsEntry", v)
}

// DecodePathSwitchRequestAcknowledgeTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePathSwitchRequestAcknowledgeTransferExtIEsEntry(data []byte) (PathSwitchRequestAcknowledgeTransferExtIEsEntry, int, error) {
	var v PathSwitchRequestAcknowledgeTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("PathSwitchRequestAcknowledgeTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PathSwitchRequestAcknowledgeTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
