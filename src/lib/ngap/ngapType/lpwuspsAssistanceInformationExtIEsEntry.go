// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type LPWUSPSAssistanceInformationExtIEsEntryPresent uint32

const (
	LPWUSPSAssistanceInformationExtIEsEntryPresentNothing LPWUSPSAssistanceInformationExtIEsEntryPresent = iota
	LPWUSPSAssistanceInformationExtIEsEntryPresentUnknown
)

type LPWUSPSAssistanceInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      LPWUSPSAssistanceInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LPWUSPSAssistanceInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LPWUSPSAssistanceInformation-ExtIEsEntry", v)
}

// DecodeLPWUSPSAssistanceInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeLPWUSPSAssistanceInformationExtIEsEntry(data []byte) (LPWUSPSAssistanceInformationExtIEsEntry, int, error) {
	var v LPWUSPSAssistanceInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("LPWUSPSAssistanceInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return LPWUSPSAssistanceInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
