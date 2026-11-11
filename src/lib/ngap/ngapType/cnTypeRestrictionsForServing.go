// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CNTypeRestrictionsForServing represents ASN.1 CNTypeRestrictionsForServing.
const (
	CNTypeRestrictionsForServingPresentEpcForbidden Enumerated = 0
)

type CNTypeRestrictionsForServing struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CNTypeRestrictionsForServing) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CNTypeRestrictionsForServing", v)
}

// DecodeCNTypeRestrictionsForServing decodes a standalone APER value and reports the consumed octets.
func DecodeCNTypeRestrictionsForServing(data []byte) (CNTypeRestrictionsForServing, int, error) {
	var v CNTypeRestrictionsForServing
	err := ngapcodec.DecodeParameter("CNTypeRestrictionsForServing", data, &v)
	if ngapcodec.IsFatal(err) {
		return CNTypeRestrictionsForServing{}, 0, err
	}
	return v, len(data), err
}
