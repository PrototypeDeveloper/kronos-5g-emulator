// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RATRestrictionInformation represents ASN.1 RATRestrictionInformation.
type RATRestrictionInformation struct {
	Value BitString `aper:"sizeExt,sizeLB:8,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RATRestrictionInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RATRestrictionInformation", v)
}

// DecodeRATRestrictionInformation decodes a standalone APER value and reports the consumed octets.
func DecodeRATRestrictionInformation(data []byte) (RATRestrictionInformation, int, error) {
	var v RATRestrictionInformation
	err := ngapcodec.DecodeParameter("RATRestrictionInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return RATRestrictionInformation{}, 0, err
	}
	return v, len(data), err
}
