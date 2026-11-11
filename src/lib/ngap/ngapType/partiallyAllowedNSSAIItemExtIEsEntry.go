// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PartiallyAllowedNSSAIItemExtIEsEntryPresent uint32

const (
	PartiallyAllowedNSSAIItemExtIEsEntryPresentNothing PartiallyAllowedNSSAIItemExtIEsEntryPresent = iota
	PartiallyAllowedNSSAIItemExtIEsEntryPresentUnknown
)

type PartiallyAllowedNSSAIItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PartiallyAllowedNSSAIItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PartiallyAllowedNSSAIItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PartiallyAllowedNSSAI-Item-ExtIEsEntry", v)
}

// DecodePartiallyAllowedNSSAIItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePartiallyAllowedNSSAIItemExtIEsEntry(data []byte) (PartiallyAllowedNSSAIItemExtIEsEntry, int, error) {
	var v PartiallyAllowedNSSAIItemExtIEsEntry
	err := ngapcodec.DecodeParameter("PartiallyAllowedNSSAI-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PartiallyAllowedNSSAIItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
