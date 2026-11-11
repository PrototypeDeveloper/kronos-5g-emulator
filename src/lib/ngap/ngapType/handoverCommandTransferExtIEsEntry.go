// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type HandoverCommandTransferExtIEsEntryPresent uint32

const (
	HandoverCommandTransferExtIEsEntryPresentNothing HandoverCommandTransferExtIEsEntryPresent = iota
	HandoverCommandTransferExtIEsEntryPresentAdditionalDLForwardingUPTNLInformation
	HandoverCommandTransferExtIEsEntryPresentULForwardingUPTNLInformation
	HandoverCommandTransferExtIEsEntryPresentAdditionalULForwardingUPTNLInformation
	HandoverCommandTransferExtIEsEntryPresentDataForwardingResponseERABList
	HandoverCommandTransferExtIEsEntryPresentQosFlowFailedToSetupList
	HandoverCommandTransferExtIEsEntryPresentUnknown
)

type HandoverCommandTransferExtIEsEntry struct {
	ID                                     ProtocolExtensionID
	Criticality                            Criticality
	Present                                HandoverCommandTransferExtIEsEntryPresent
	AdditionalDLForwardingUPTNLInformation QosFlowPerTNLInformationList    `ngap:"id:152,criticality:ignore,presence:optional"`
	ULForwardingUPTNLInformation           UPTransportLayerInformation     `ngap:"id:164,criticality:reject,presence:optional"`
	AdditionalULForwardingUPTNLInformation UPTransportLayerInformationList `ngap:"id:172,criticality:reject,presence:optional"`
	DataForwardingResponseERABList         DataForwardingResponseERABList  `ngap:"id:249,criticality:ignore,presence:optional"`
	QosFlowFailedToSetupList               QosFlowListWithCause            `ngap:"id:283,criticality:ignore,presence:optional"`
	UnknownValue                           []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HandoverCommandTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("HandoverCommandTransfer-ExtIEsEntry", v)
}

// DecodeHandoverCommandTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeHandoverCommandTransferExtIEsEntry(data []byte) (HandoverCommandTransferExtIEsEntry, int, error) {
	var v HandoverCommandTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("HandoverCommandTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return HandoverCommandTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
