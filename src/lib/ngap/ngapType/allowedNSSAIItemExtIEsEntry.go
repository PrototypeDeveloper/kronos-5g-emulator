// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AllowedNSSAIItemExtIEsEntryPresent uint32

const (
	AllowedNSSAIItemExtIEsEntryPresentNothing AllowedNSSAIItemExtIEsEntryPresent = iota
	AllowedNSSAIItemExtIEsEntryPresentUnknown
)

type AllowedNSSAIItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AllowedNSSAIItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AllowedNSSAIItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AllowedNSSAI-Item-ExtIEsEntry", v)
}

// DecodeAllowedNSSAIItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAllowedNSSAIItemExtIEsEntry(data []byte) (AllowedNSSAIItemExtIEsEntry, int, error) {
	var v AllowedNSSAIItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AllowedNSSAI-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AllowedNSSAIItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
