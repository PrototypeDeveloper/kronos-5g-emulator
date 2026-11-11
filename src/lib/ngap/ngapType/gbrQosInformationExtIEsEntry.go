// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type GBRQosInformationExtIEsEntryPresent uint32

const (
	GBRQosInformationExtIEsEntryPresentNothing GBRQosInformationExtIEsEntryPresent = iota
	GBRQosInformationExtIEsEntryPresentAlternativeQoSParaSetList
	GBRQosInformationExtIEsEntryPresentMonitoringRequestonAvailableBitrate
	GBRQosInformationExtIEsEntryPresentUnknown
)

type GBRQosInformationExtIEsEntry struct {
	ID                                  ProtocolExtensionID
	Criticality                         Criticality
	Present                             GBRQosInformationExtIEsEntryPresent
	AlternativeQoSParaSetList           AlternativeQoSParaSetList           `ngap:"id:220,criticality:ignore,presence:optional"`
	MonitoringRequestonAvailableBitrate MonitoringRequestonAvailableBitrate `ngap:"id:489,criticality:ignore,presence:optional"`
	UnknownValue                        []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GBRQosInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GBR-QosInformation-ExtIEsEntry", v)
}

// DecodeGBRQosInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeGBRQosInformationExtIEsEntry(data []byte) (GBRQosInformationExtIEsEntry, int, error) {
	var v GBRQosInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("GBR-QosInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return GBRQosInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
