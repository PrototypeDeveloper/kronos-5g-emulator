// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TNLAddressWeightFactor represents ASN.1 TNLAddressWeightFactor.
type TNLAddressWeightFactor struct {
	Value int64 `aper:"valueLB:0,valueUB:255"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TNLAddressWeightFactor) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TNLAddressWeightFactor", v)
}

// DecodeTNLAddressWeightFactor decodes a standalone APER value and reports the consumed octets.
func DecodeTNLAddressWeightFactor(data []byte) (TNLAddressWeightFactor, int, error) {
	var v TNLAddressWeightFactor
	err := ngapcodec.DecodeParameter("TNLAddressWeightFactor", data, &v)
	if ngapcodec.IsFatal(err) {
		return TNLAddressWeightFactor{}, 0, err
	}
	return v, len(data), err
}
