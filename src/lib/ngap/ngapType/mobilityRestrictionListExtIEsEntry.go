// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MobilityRestrictionListExtIEsEntryPresent uint32

const (
	MobilityRestrictionListExtIEsEntryPresentNothing MobilityRestrictionListExtIEsEntryPresent = iota
	MobilityRestrictionListExtIEsEntryPresentLastEUTRANPLMNIdentity
	MobilityRestrictionListExtIEsEntryPresentCNTypeRestrictionsForServing
	MobilityRestrictionListExtIEsEntryPresentCNTypeRestrictionsForEquivalent
	MobilityRestrictionListExtIEsEntryPresentNPNMobilityInformation
	MobilityRestrictionListExtIEsEntryPresentUnknown
)

type MobilityRestrictionListExtIEsEntry struct {
	ID                              ProtocolExtensionID
	Criticality                     Criticality
	Present                         MobilityRestrictionListExtIEsEntryPresent
	LastEUTRANPLMNIdentity          PLMNIdentity                    `ngap:"id:150,criticality:ignore,presence:optional"`
	CNTypeRestrictionsForServing    CNTypeRestrictionsForServing    `ngap:"id:161,criticality:ignore,presence:optional"`
	CNTypeRestrictionsForEquivalent CNTypeRestrictionsForEquivalent `ngap:"id:160,criticality:ignore,presence:optional"`
	NPNMobilityInformation          NPNMobilityInformation          `ngap:"id:261,criticality:reject,presence:optional"`
	UnknownValue                    []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MobilityRestrictionListExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MobilityRestrictionList-ExtIEsEntry", v)
}

// DecodeMobilityRestrictionListExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMobilityRestrictionListExtIEsEntry(data []byte) (MobilityRestrictionListExtIEsEntry, int, error) {
	var v MobilityRestrictionListExtIEsEntry
	err := ngapcodec.DecodeParameter("MobilityRestrictionList-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MobilityRestrictionListExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
