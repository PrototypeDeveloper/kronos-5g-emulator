// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TargetNSSAIItemExtIEsEntryPresent uint32

const (
	TargetNSSAIItemExtIEsEntryPresentNothing TargetNSSAIItemExtIEsEntryPresent = iota
	TargetNSSAIItemExtIEsEntryPresentUnknown
)

type TargetNSSAIItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TargetNSSAIItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetNSSAIItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetNSSAI-Item-ExtIEsEntry", v)
}

// DecodeTargetNSSAIItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTargetNSSAIItemExtIEsEntry(data []byte) (TargetNSSAIItemExtIEsEntry, int, error) {
	var v TargetNSSAIItemExtIEsEntry
	err := ngapcodec.DecodeParameter("TargetNSSAI-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetNSSAIItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
