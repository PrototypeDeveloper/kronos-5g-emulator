// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type WUSAssistanceInformationExtIEsEntryPresent uint32

const (
	WUSAssistanceInformationExtIEsEntryPresentNothing WUSAssistanceInformationExtIEsEntryPresent = iota
	WUSAssistanceInformationExtIEsEntryPresentUnknown
)

type WUSAssistanceInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      WUSAssistanceInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v WUSAssistanceInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("WUS-Assistance-Information-ExtIEsEntry", v)
}

// DecodeWUSAssistanceInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeWUSAssistanceInformationExtIEsEntry(data []byte) (WUSAssistanceInformationExtIEsEntry, int, error) {
	var v WUSAssistanceInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("WUS-Assistance-Information-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return WUSAssistanceInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
