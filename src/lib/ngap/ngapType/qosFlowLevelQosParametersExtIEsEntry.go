// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QosFlowLevelQosParametersExtIEsEntryPresent uint32

const (
	QosFlowLevelQosParametersExtIEsEntryPresentNothing QosFlowLevelQosParametersExtIEsEntryPresent = iota
	QosFlowLevelQosParametersExtIEsEntryPresentQosMonitoringRequest
	QosFlowLevelQosParametersExtIEsEntryPresentQosMonitoringReportingFrequency
	QosFlowLevelQosParametersExtIEsEntryPresentPDUsetQoSParameters
	QosFlowLevelQosParametersExtIEsEntryPresentDLPDUSetInformationMarkingSupportIndication
	QosFlowLevelQosParametersExtIEsEntryPresentMMSID
	QosFlowLevelQosParametersExtIEsEntryPresentIndicationOfBitrateAdaptation
	QosFlowLevelQosParametersExtIEsEntryPresentUnknown
)

type QosFlowLevelQosParametersExtIEsEntry struct {
	ID                                          ProtocolExtensionID
	Criticality                                 Criticality
	Present                                     QosFlowLevelQosParametersExtIEsEntryPresent
	QosMonitoringRequest                        QosMonitoringRequest                        `ngap:"id:181,criticality:ignore,presence:optional"`
	QosMonitoringReportingFrequency             QosMonitoringReportingFrequency             `ngap:"id:276,criticality:ignore,presence:optional"`
	PDUsetQoSParameters                         PDUsetQoSParameters                         `ngap:"id:422,criticality:ignore,presence:optional"`
	DLPDUSetInformationMarkingSupportIndication DLPDUSetInformationMarkingSupportIndication `ngap:"id:488,criticality:ignore,presence:optional"`
	MMSID                                       MMSID                                       `ngap:"id:490,criticality:ignore,presence:optional"`
	IndicationOfBitrateAdaptation               IndicationOfBitrateAdaptation               `ngap:"id:491,criticality:ignore,presence:optional"`
	UnknownValue                                []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowLevelQosParametersExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowLevelQosParameters-ExtIEsEntry", v)
}

// DecodeQosFlowLevelQosParametersExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowLevelQosParametersExtIEsEntry(data []byte) (QosFlowLevelQosParametersExtIEsEntry, int, error) {
	var v QosFlowLevelQosParametersExtIEsEntry
	err := ngapcodec.DecodeParameter("QosFlowLevelQosParameters-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowLevelQosParametersExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
