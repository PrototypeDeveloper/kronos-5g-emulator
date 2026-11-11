// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TNGFIdentityInformation represents ASN.1 TNGFIdentityInformation.
type TNGFIdentityInformation struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TNGFIdentityInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TNGFIdentityInformation", v)
}

// DecodeTNGFIdentityInformation decodes a standalone APER value and reports the consumed octets.
func DecodeTNGFIdentityInformation(data []byte) (TNGFIdentityInformation, int, error) {
	var v TNGFIdentityInformation
	err := ngapcodec.DecodeParameter("TNGFIdentityInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return TNGFIdentityInformation{}, 0, err
	}
	return v, len(data), err
}
