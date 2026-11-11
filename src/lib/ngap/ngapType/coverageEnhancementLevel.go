// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CoverageEnhancementLevel represents ASN.1 CoverageEnhancementLevel.
type CoverageEnhancementLevel struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CoverageEnhancementLevel) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CoverageEnhancementLevel", v)
}

// DecodeCoverageEnhancementLevel decodes a standalone APER value and reports the consumed octets.
func DecodeCoverageEnhancementLevel(data []byte) (CoverageEnhancementLevel, int, error) {
	var v CoverageEnhancementLevel
	err := ngapcodec.DecodeParameter("CoverageEnhancementLevel", data, &v)
	if ngapcodec.IsFatal(err) {
		return CoverageEnhancementLevel{}, 0, err
	}
	return v, len(data), err
}
