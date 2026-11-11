// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RATRestrictionsItem represents ASN.1 RATRestrictions-Item.
type RATRestrictionsItem struct {
	PLMNIdentity              PLMNIdentity
	RATRestrictionInformation RATRestrictionInformation
	HasIEExtensions           bool
	IEExtensions              []RATRestrictionsItemExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RATRestrictionsItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RATRestrictions-Item", v)
}

// DecodeRATRestrictionsItem decodes a standalone APER value and reports the consumed octets.
func DecodeRATRestrictionsItem(data []byte) (RATRestrictionsItem, int, error) {
	var v RATRestrictionsItem
	err := ngapcodec.DecodeParameter("RATRestrictions-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return RATRestrictionsItem{}, 0, err
	}
	return v, len(data), err
}
