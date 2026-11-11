// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NRNTNTAIInformationExtIEsEntryPresent uint32

const (
	NRNTNTAIInformationExtIEsEntryPresentNothing NRNTNTAIInformationExtIEsEntryPresent = iota
	NRNTNTAIInformationExtIEsEntryPresentUnknown
)

type NRNTNTAIInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NRNTNTAIInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRNTNTAIInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NRNTNTAIInformation-ExtIEsEntry", v)
}

// DecodeNRNTNTAIInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNRNTNTAIInformationExtIEsEntry(data []byte) (NRNTNTAIInformationExtIEsEntry, int, error) {
	var v NRNTNTAIInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("NRNTNTAIInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRNTNTAIInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
