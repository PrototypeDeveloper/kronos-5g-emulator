// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceModifyRequestTransferIEsEntryPresent uint32

const (
	PDUSessionResourceModifyRequestTransferIEsEntryPresentNothing PDUSessionResourceModifyRequestTransferIEsEntryPresent = iota
	PDUSessionResourceModifyRequestTransferIEsEntryPresentPDUSessionAggregateMaximumBitRate
	PDUSessionResourceModifyRequestTransferIEsEntryPresentULNGUUPTNLModifyList
	PDUSessionResourceModifyRequestTransferIEsEntryPresentNetworkInstance
	PDUSessionResourceModifyRequestTransferIEsEntryPresentQosFlowAddOrModifyRequestList
	PDUSessionResourceModifyRequestTransferIEsEntryPresentQosFlowToReleaseList
	PDUSessionResourceModifyRequestTransferIEsEntryPresentAdditionalULNGUUPTNLInformation
	PDUSessionResourceModifyRequestTransferIEsEntryPresentCommonNetworkInstance
	PDUSessionResourceModifyRequestTransferIEsEntryPresentAdditionalRedundantULNGUUPTNLInformation
	PDUSessionResourceModifyRequestTransferIEsEntryPresentRedundantCommonNetworkInstance
	PDUSessionResourceModifyRequestTransferIEsEntryPresentRedundantULNGUUPTNLInformation
	PDUSessionResourceModifyRequestTransferIEsEntryPresentSecurityIndication
	PDUSessionResourceModifyRequestTransferIEsEntryPresentMBSSessionSetuporModifyRequestList
	PDUSessionResourceModifyRequestTransferIEsEntryPresentMBSSessionToReleaseList
	PDUSessionResourceModifyRequestTransferIEsEntryPresentUserPlaneFailureIndication
	PDUSessionResourceModifyRequestTransferIEsEntryPresentUnknown
)

type PDUSessionResourceModifyRequestTransferIEsEntry struct {
	ID                                       ProtocolIEID
	Criticality                              Criticality
	Present                                  PDUSessionResourceModifyRequestTransferIEsEntryPresent
	PDUSessionAggregateMaximumBitRate        PDUSessionAggregateMaximumBitRate  `ngap:"id:130,criticality:reject,presence:optional"`
	ULNGUUPTNLModifyList                     ULNGUUPTNLModifyList               `ngap:"id:140,criticality:reject,presence:optional"`
	NetworkInstance                          NetworkInstance                    `ngap:"id:129,criticality:reject,presence:optional"`
	QosFlowAddOrModifyRequestList            QosFlowAddOrModifyRequestList      `ngap:"id:135,criticality:reject,presence:optional"`
	QosFlowToReleaseList                     QosFlowListWithCause               `ngap:"id:137,criticality:reject,presence:optional"`
	AdditionalULNGUUPTNLInformation          UPTransportLayerInformationList    `ngap:"id:126,criticality:reject,presence:optional"`
	CommonNetworkInstance                    CommonNetworkInstance              `ngap:"id:166,criticality:ignore,presence:optional"`
	AdditionalRedundantULNGUUPTNLInformation UPTransportLayerInformationList    `ngap:"id:186,criticality:ignore,presence:optional"`
	RedundantCommonNetworkInstance           CommonNetworkInstance              `ngap:"id:190,criticality:ignore,presence:optional"`
	RedundantULNGUUPTNLInformation           UPTransportLayerInformation        `ngap:"id:195,criticality:ignore,presence:optional"`
	SecurityIndication                       SecurityIndication                 `ngap:"id:138,criticality:ignore,presence:optional"`
	MBSSessionSetuporModifyRequestList       MBSSessionSetuporModifyRequestList `ngap:"id:319,criticality:ignore,presence:optional"`
	MBSSessionToReleaseList                  MBSSessionToReleaseList            `ngap:"id:317,criticality:ignore,presence:optional"`
	UserPlaneFailureIndication               UserPlaneFailureIndication         `ngap:"id:435,criticality:ignore,presence:optional"`
	UnknownValue                             []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyRequestTransferIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyRequestTransferIEsEntry", v)
}

// DecodePDUSessionResourceModifyRequestTransferIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyRequestTransferIEsEntry(data []byte) (PDUSessionResourceModifyRequestTransferIEsEntry, int, error) {
	var v PDUSessionResourceModifyRequestTransferIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyRequestTransferIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyRequestTransferIEsEntry{}, 0, err
	}
	return v, len(data), err
}
