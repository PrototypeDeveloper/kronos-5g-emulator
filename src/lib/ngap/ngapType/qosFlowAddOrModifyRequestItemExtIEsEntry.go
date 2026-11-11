// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QosFlowAddOrModifyRequestItemExtIEsEntryPresent uint32

const (
	QosFlowAddOrModifyRequestItemExtIEsEntryPresentNothing QosFlowAddOrModifyRequestItemExtIEsEntryPresent = iota
	QosFlowAddOrModifyRequestItemExtIEsEntryPresentTSCTrafficCharacteristics
	QosFlowAddOrModifyRequestItemExtIEsEntryPresentRedundantQosFlowIndicator
	QosFlowAddOrModifyRequestItemExtIEsEntryPresentULNGUUPTNLInformation
	QosFlowAddOrModifyRequestItemExtIEsEntryPresentDownlinkTLContainer
	QosFlowAddOrModifyRequestItemExtIEsEntryPresentECNMarkingorCongestionInformationReportingRequest
	QosFlowAddOrModifyRequestItemExtIEsEntryPresentUnknown
)

type QosFlowAddOrModifyRequestItemExtIEsEntry struct {
	ID                                                ProtocolExtensionID
	Criticality                                       Criticality
	Present                                           QosFlowAddOrModifyRequestItemExtIEsEntryPresent
	TSCTrafficCharacteristics                         TSCTrafficCharacteristics                         `ngap:"id:196,criticality:ignore,presence:optional"`
	RedundantQosFlowIndicator                         RedundantQosFlowIndicator                         `ngap:"id:194,criticality:ignore,presence:optional"`
	ULNGUUPTNLInformation                             UPTransportLayerInformation                       `ngap:"id:139,criticality:ignore,presence:optional"`
	DownlinkTLContainer                               TLContainer                                       `ngap:"id:394,criticality:ignore,presence:optional"`
	ECNMarkingorCongestionInformationReportingRequest ECNMarkingorCongestionInformationReportingRequest `ngap:"id:425,criticality:ignore,presence:optional"`
	UnknownValue                                      []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowAddOrModifyRequestItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowAddOrModifyRequestItem-ExtIEsEntry", v)
}

// DecodeQosFlowAddOrModifyRequestItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowAddOrModifyRequestItemExtIEsEntry(data []byte) (QosFlowAddOrModifyRequestItemExtIEsEntry, int, error) {
	var v QosFlowAddOrModifyRequestItemExtIEsEntry
	err := ngapcodec.DecodeParameter("QosFlowAddOrModifyRequestItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowAddOrModifyRequestItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
