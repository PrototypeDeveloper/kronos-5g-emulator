// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TSCAssistanceInformationExtIEsEntryPresent uint32

const (
	TSCAssistanceInformationExtIEsEntryPresentNothing TSCAssistanceInformationExtIEsEntryPresent = iota
	TSCAssistanceInformationExtIEsEntryPresentSurvivalTime
	TSCAssistanceInformationExtIEsEntryPresentRANfeedbacktype
	TSCAssistanceInformationExtIEsEntryPresentN6JitterInformation
	TSCAssistanceInformationExtIEsEntryPresentUnknown
)

type TSCAssistanceInformationExtIEsEntry struct {
	ID                  ProtocolExtensionID
	Criticality         Criticality
	Present             TSCAssistanceInformationExtIEsEntryPresent
	SurvivalTime        SurvivalTime        `ngap:"id:327,criticality:ignore,presence:optional"`
	RANfeedbacktype     RANfeedbacktype     `ngap:"id:391,criticality:ignore,presence:optional"`
	N6JitterInformation N6JitterInformation `ngap:"id:424,criticality:ignore,presence:optional"`
	UnknownValue        []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TSCAssistanceInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TSCAssistanceInformation-ExtIEsEntry", v)
}

// DecodeTSCAssistanceInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTSCAssistanceInformationExtIEsEntry(data []byte) (TSCAssistanceInformationExtIEsEntry, int, error) {
	var v TSCAssistanceInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("TSCAssistanceInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TSCAssistanceInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
