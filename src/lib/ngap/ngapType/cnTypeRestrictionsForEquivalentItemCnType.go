// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CNTypeRestrictionsForEquivalentItemCnType represents ASN.1 CNTypeRestrictionsForEquivalentItemCnType.
const (
	CNTypeRestrictionsForEquivalentItemCnTypePresentEpcForbidden    Enumerated = 0
	CNTypeRestrictionsForEquivalentItemCnTypePresentFiveGCForbidden Enumerated = 1
)

type CNTypeRestrictionsForEquivalentItemCnType struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v CNTypeRestrictionsForEquivalentItemCnType) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:CNTypeRestrictionsForEquivalentItemCnType", v)
}

// DecodeCNTypeRestrictionsForEquivalentItemCnType decodes a standalone APER value and reports consumed octets.
func DecodeCNTypeRestrictionsForEquivalentItemCnType(data []byte) (CNTypeRestrictionsForEquivalentItemCnType, int, error) {
	var v CNTypeRestrictionsForEquivalentItemCnType
	err := ngapcodec.DecodeParameter("inline:CNTypeRestrictionsForEquivalentItemCnType", data, &v)
	if ngapcodec.IsFatal(err) {
		return CNTypeRestrictionsForEquivalentItemCnType{}, 0, err
	}
	return v, len(data), err
}
