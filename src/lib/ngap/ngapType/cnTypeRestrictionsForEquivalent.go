// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CNTypeRestrictionsForEquivalent represents ASN.1 CNTypeRestrictionsForEquivalent.
type CNTypeRestrictionsForEquivalent struct {
	List []CNTypeRestrictionsForEquivalentItem `aper:"sizeLB:1,sizeUB:15"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CNTypeRestrictionsForEquivalent) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CNTypeRestrictionsForEquivalent", v)
}

// DecodeCNTypeRestrictionsForEquivalent decodes a standalone APER value and reports the consumed octets.
func DecodeCNTypeRestrictionsForEquivalent(data []byte) (CNTypeRestrictionsForEquivalent, int, error) {
	var v CNTypeRestrictionsForEquivalent
	err := ngapcodec.DecodeParameter("CNTypeRestrictionsForEquivalent", data, &v)
	if ngapcodec.IsFatal(err) {
		return CNTypeRestrictionsForEquivalent{}, 0, err
	}
	return v, len(data), err
}
