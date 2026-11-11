// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SupportedTAItemExtIEsEntryPresent uint32

const (
	SupportedTAItemExtIEsEntryPresentNothing SupportedTAItemExtIEsEntryPresent = iota
	SupportedTAItemExtIEsEntryPresentConfiguredTACIndication
	SupportedTAItemExtIEsEntryPresentRATInformation
	SupportedTAItemExtIEsEntryPresentUnknown
)

type SupportedTAItemExtIEsEntry struct {
	ID                      ProtocolExtensionID
	Criticality             Criticality
	Present                 SupportedTAItemExtIEsEntryPresent
	ConfiguredTACIndication ConfiguredTACIndication `ngap:"id:272,criticality:ignore,presence:optional"`
	RATInformation          RATInformation          `ngap:"id:179,criticality:reject,presence:optional"`
	UnknownValue            []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SupportedTAItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SupportedTAItem-ExtIEsEntry", v)
}

// DecodeSupportedTAItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSupportedTAItemExtIEsEntry(data []byte) (SupportedTAItemExtIEsEntry, int, error) {
	var v SupportedTAItemExtIEsEntry
	err := ngapcodec.DecodeParameter("SupportedTAItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SupportedTAItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
