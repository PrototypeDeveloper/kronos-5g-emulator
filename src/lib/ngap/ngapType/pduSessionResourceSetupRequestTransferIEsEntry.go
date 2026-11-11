// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceSetupRequestTransferIEsEntryPresent uint32

const (
	PDUSessionResourceSetupRequestTransferIEsEntryPresentNothing PDUSessionResourceSetupRequestTransferIEsEntryPresent = iota
	PDUSessionResourceSetupRequestTransferIEsEntryPresentPDUSessionAggregateMaximumBitRate
	PDUSessionResourceSetupRequestTransferIEsEntryPresentULNGUUPTNLInformation
	PDUSessionResourceSetupRequestTransferIEsEntryPresentAdditionalULNGUUPTNLInformation
	PDUSessionResourceSetupRequestTransferIEsEntryPresentDataForwardingNotPossible
	PDUSessionResourceSetupRequestTransferIEsEntryPresentPDUSessionType
	PDUSessionResourceSetupRequestTransferIEsEntryPresentSecurityIndication
	PDUSessionResourceSetupRequestTransferIEsEntryPresentNetworkInstance
	PDUSessionResourceSetupRequestTransferIEsEntryPresentQosFlowSetupRequestList
	PDUSessionResourceSetupRequestTransferIEsEntryPresentCommonNetworkInstance
	PDUSessionResourceSetupRequestTransferIEsEntryPresentDirectForwardingPathAvailability
	PDUSessionResourceSetupRequestTransferIEsEntryPresentRedundantULNGUUPTNLInformation
	PDUSessionResourceSetupRequestTransferIEsEntryPresentAdditionalRedundantULNGUUPTNLInformation
	PDUSessionResourceSetupRequestTransferIEsEntryPresentRedundantCommonNetworkInstance
	PDUSessionResourceSetupRequestTransferIEsEntryPresentRedundantPDUSessionInformation
	PDUSessionResourceSetupRequestTransferIEsEntryPresentMBSSessionSetupRequestList
	PDUSessionResourceSetupRequestTransferIEsEntryPresentDownlinkTLContainer
	PDUSessionResourceSetupRequestTransferIEsEntryPresentUnknown
)

type PDUSessionResourceSetupRequestTransferIEsEntry struct {
	ID                                       ProtocolIEID
	Criticality                              Criticality
	Present                                  PDUSessionResourceSetupRequestTransferIEsEntryPresent
	PDUSessionAggregateMaximumBitRate        PDUSessionAggregateMaximumBitRate `ngap:"id:130,criticality:reject,presence:optional"`
	ULNGUUPTNLInformation                    UPTransportLayerInformation       `ngap:"id:139,criticality:reject,presence:mandatory"`
	AdditionalULNGUUPTNLInformation          UPTransportLayerInformationList   `ngap:"id:126,criticality:reject,presence:optional"`
	DataForwardingNotPossible                DataForwardingNotPossible         `ngap:"id:127,criticality:reject,presence:optional"`
	PDUSessionType                           PDUSessionType                    `ngap:"id:134,criticality:reject,presence:mandatory"`
	SecurityIndication                       SecurityIndication                `ngap:"id:138,criticality:reject,presence:optional"`
	NetworkInstance                          NetworkInstance                   `ngap:"id:129,criticality:reject,presence:optional"`
	QosFlowSetupRequestList                  QosFlowSetupRequestList           `ngap:"id:136,criticality:reject,presence:mandatory"`
	CommonNetworkInstance                    CommonNetworkInstance             `ngap:"id:166,criticality:ignore,presence:optional"`
	DirectForwardingPathAvailability         DirectForwardingPathAvailability  `ngap:"id:22,criticality:ignore,presence:optional"`
	RedundantULNGUUPTNLInformation           UPTransportLayerInformation       `ngap:"id:195,criticality:ignore,presence:optional"`
	AdditionalRedundantULNGUUPTNLInformation UPTransportLayerInformationList   `ngap:"id:186,criticality:ignore,presence:optional"`
	RedundantCommonNetworkInstance           CommonNetworkInstance             `ngap:"id:190,criticality:ignore,presence:optional"`
	RedundantPDUSessionInformation           RedundantPDUSessionInformation    `ngap:"id:197,criticality:ignore,presence:optional"`
	MBSSessionSetupRequestList               MBSSessionSetupRequestList        `ngap:"id:318,criticality:ignore,presence:optional"`
	DownlinkTLContainer                      TLContainer                       `ngap:"id:394,criticality:ignore,presence:optional"`
	UnknownValue                             []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupRequestTransferIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupRequestTransferIEsEntry", v)
}

// DecodePDUSessionResourceSetupRequestTransferIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupRequestTransferIEsEntry(data []byte) (PDUSessionResourceSetupRequestTransferIEsEntry, int, error) {
	var v PDUSessionResourceSetupRequestTransferIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupRequestTransferIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupRequestTransferIEsEntry{}, 0, err
	}
	return v, len(data), err
}
