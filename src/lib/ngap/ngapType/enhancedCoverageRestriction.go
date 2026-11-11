// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EnhancedCoverageRestriction represents ASN.1 Enhanced-CoverageRestriction.
const (
	EnhancedCoverageRestrictionPresentRestricted Enumerated = 0
)

type EnhancedCoverageRestriction struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EnhancedCoverageRestriction) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Enhanced-CoverageRestriction", v)
}

// DecodeEnhancedCoverageRestriction decodes a standalone APER value and reports the consumed octets.
func DecodeEnhancedCoverageRestriction(data []byte) (EnhancedCoverageRestriction, int, error) {
	var v EnhancedCoverageRestriction
	err := ngapcodec.DecodeParameter("Enhanced-CoverageRestriction", data, &v)
	if ngapcodec.IsFatal(err) {
		return EnhancedCoverageRestriction{}, 0, err
	}
	return v, len(data), err
}
