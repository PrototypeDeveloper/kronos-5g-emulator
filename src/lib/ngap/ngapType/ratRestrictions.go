// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RATRestrictions represents ASN.1 RATRestrictions.
type RATRestrictions struct {
	List []RATRestrictionsItem `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RATRestrictions) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RATRestrictions", v)
}

// DecodeRATRestrictions decodes a standalone APER value and reports the consumed octets.
func DecodeRATRestrictions(data []byte) (RATRestrictions, int, error) {
	var v RATRestrictions
	err := ngapcodec.DecodeParameter("RATRestrictions", data, &v)
	if ngapcodec.IsFatal(err) {
		return RATRestrictions{}, 0, err
	}
	return v, len(data), err
}
