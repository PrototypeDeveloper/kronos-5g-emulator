// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AIoTCommandAssistanceInformationExtIEsEntryPresent uint32

const (
	AIoTCommandAssistanceInformationExtIEsEntryPresentNothing AIoTCommandAssistanceInformationExtIEsEntryPresent = iota
	AIoTCommandAssistanceInformationExtIEsEntryPresentUnknown
)

type AIoTCommandAssistanceInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AIoTCommandAssistanceInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTCommandAssistanceInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-CommandAssistanceInformation-ExtIEsEntry", v)
}

// DecodeAIoTCommandAssistanceInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTCommandAssistanceInformationExtIEsEntry(data []byte) (AIoTCommandAssistanceInformationExtIEsEntry, int, error) {
	var v AIoTCommandAssistanceInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("AIoT-CommandAssistanceInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTCommandAssistanceInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
