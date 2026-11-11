// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QosFlowSetupRequestItemExtIEsEntryPresent uint32

const (
	QosFlowSetupRequestItemExtIEsEntryPresentNothing QosFlowSetupRequestItemExtIEsEntryPresent = iota
	QosFlowSetupRequestItemExtIEsEntryPresentTSCTrafficCharacteristics
	QosFlowSetupRequestItemExtIEsEntryPresentRedundantQosFlowIndicator
	QosFlowSetupRequestItemExtIEsEntryPresentECNMarkingorCongestionInformationReportingRequest
	QosFlowSetupRequestItemExtIEsEntryPresentUnknown
)

type QosFlowSetupRequestItemExtIEsEntry struct {
	ID                                                ProtocolExtensionID
	Criticality                                       Criticality
	Present                                           QosFlowSetupRequestItemExtIEsEntryPresent
	TSCTrafficCharacteristics                         TSCTrafficCharacteristics                         `ngap:"id:196,criticality:ignore,presence:optional"`
	RedundantQosFlowIndicator                         RedundantQosFlowIndicator                         `ngap:"id:194,criticality:ignore,presence:optional"`
	ECNMarkingorCongestionInformationReportingRequest ECNMarkingorCongestionInformationReportingRequest `ngap:"id:425,criticality:ignore,presence:optional"`
	UnknownValue                                      []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowSetupRequestItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowSetupRequestItem-ExtIEsEntry", v)
}

// DecodeQosFlowSetupRequestItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowSetupRequestItemExtIEsEntry(data []byte) (QosFlowSetupRequestItemExtIEsEntry, int, error) {
	var v QosFlowSetupRequestItemExtIEsEntry
	err := ngapcodec.DecodeParameter("QosFlowSetupRequestItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowSetupRequestItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
