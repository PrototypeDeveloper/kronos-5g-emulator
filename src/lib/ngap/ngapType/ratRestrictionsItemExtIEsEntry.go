// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RATRestrictionsItemExtIEsEntryPresent uint32

const (
	RATRestrictionsItemExtIEsEntryPresentNothing RATRestrictionsItemExtIEsEntryPresent = iota
	RATRestrictionsItemExtIEsEntryPresentExtendedRATRestrictionInformation
	RATRestrictionsItemExtIEsEntryPresentUnknown
)

type RATRestrictionsItemExtIEsEntry struct {
	ID                                ProtocolExtensionID
	Criticality                       Criticality
	Present                           RATRestrictionsItemExtIEsEntryPresent
	ExtendedRATRestrictionInformation ExtendedRATRestrictionInformation `ngap:"id:180,criticality:ignore,presence:optional"`
	UnknownValue                      []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RATRestrictionsItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RATRestrictions-Item-ExtIEsEntry", v)
}

// DecodeRATRestrictionsItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRATRestrictionsItemExtIEsEntry(data []byte) (RATRestrictionsItemExtIEsEntry, int, error) {
	var v RATRestrictionsItemExtIEsEntry
	err := ngapcodec.DecodeParameter("RATRestrictions-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RATRestrictionsItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
