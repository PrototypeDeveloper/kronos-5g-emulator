// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MobilityRestrictionList represents ASN.1 MobilityRestrictionList.
type MobilityRestrictionList struct {
	ServingPLMN                 PLMNIdentity
	HasEquivalentPLMNs          bool
	EquivalentPLMNs             EquivalentPLMNs
	HasRATRestrictions          bool
	RATRestrictions             RATRestrictions
	HasForbiddenAreaInformation bool
	ForbiddenAreaInformation    ForbiddenAreaInformation
	HasServiceAreaInformation   bool
	ServiceAreaInformation      ServiceAreaInformation
	HasIEExtensions             bool
	IEExtensions                []MobilityRestrictionListExtIEsEntry
	UnknownExtensions           []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MobilityRestrictionList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MobilityRestrictionList", v)
}

// DecodeMobilityRestrictionList decodes a standalone APER value and reports the consumed octets.
func DecodeMobilityRestrictionList(data []byte) (MobilityRestrictionList, int, error) {
	var v MobilityRestrictionList
	err := ngapcodec.DecodeParameter("MobilityRestrictionList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MobilityRestrictionList{}, 0, err
	}
	return v, len(data), err
}
