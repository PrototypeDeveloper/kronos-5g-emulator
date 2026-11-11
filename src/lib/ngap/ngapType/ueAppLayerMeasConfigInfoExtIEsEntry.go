// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UEAppLayerMeasConfigInfoExtIEsEntryPresent uint32

const (
	UEAppLayerMeasConfigInfoExtIEsEntryPresentNothing UEAppLayerMeasConfigInfoExtIEsEntryPresent = iota
	UEAppLayerMeasConfigInfoExtIEsEntryPresentAssistanceInformationQoEMeas
	UEAppLayerMeasConfigInfoExtIEsEntryPresentCommServiceType
	UEAppLayerMeasConfigInfoExtIEsEntryPresentQoERVQoEReportingPaths
	UEAppLayerMeasConfigInfoExtIEsEntryPresentUnknown
)

type UEAppLayerMeasConfigInfoExtIEsEntry struct {
	ID                           ProtocolExtensionID
	Criticality                  Criticality
	Present                      UEAppLayerMeasConfigInfoExtIEsEntryPresent
	AssistanceInformationQoEMeas AssistanceInformationQoEMeas `ngap:"id:398,criticality:ignore,presence:optional"`
	CommServiceType              CommServiceType              `ngap:"id:399,criticality:ignore,presence:optional"`
	QoERVQoEReportingPaths       QoERVQoEReportingPaths       `ngap:"id:438,criticality:ignore,presence:optional"`
	UnknownValue                 []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEAppLayerMeasConfigInfoExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEAppLayerMeasConfigInfo-ExtIEsEntry", v)
}

// DecodeUEAppLayerMeasConfigInfoExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUEAppLayerMeasConfigInfoExtIEsEntry(data []byte) (UEAppLayerMeasConfigInfoExtIEsEntry, int, error) {
	var v UEAppLayerMeasConfigInfoExtIEsEntry
	err := ngapcodec.DecodeParameter("UEAppLayerMeasConfigInfo-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEAppLayerMeasConfigInfoExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
