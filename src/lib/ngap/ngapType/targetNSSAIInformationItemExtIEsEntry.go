// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TargetNSSAIInformationItemExtIEsEntryPresent uint32

const (
	TargetNSSAIInformationItemExtIEsEntryPresentNothing TargetNSSAIInformationItemExtIEsEntryPresent = iota
	TargetNSSAIInformationItemExtIEsEntryPresentUnknown
)

type TargetNSSAIInformationItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TargetNSSAIInformationItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetNSSAIInformationItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetNSSAIInformation-Item-ExtIEsEntry", v)
}

// DecodeTargetNSSAIInformationItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTargetNSSAIInformationItemExtIEsEntry(data []byte) (TargetNSSAIInformationItemExtIEsEntry, int, error) {
	var v TargetNSSAIInformationItemExtIEsEntry
	err := ngapcodec.DecodeParameter("TargetNSSAIInformation-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetNSSAIInformationItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
