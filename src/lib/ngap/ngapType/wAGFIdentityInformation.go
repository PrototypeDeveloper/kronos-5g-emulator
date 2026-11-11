// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// WAGFIdentityInformation represents ASN.1 W-AGFIdentityInformation.
type WAGFIdentityInformation struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v WAGFIdentityInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("W-AGFIdentityInformation", v)
}

// DecodeWAGFIdentityInformation decodes a standalone APER value and reports the consumed octets.
func DecodeWAGFIdentityInformation(data []byte) (WAGFIdentityInformation, int, error) {
	var v WAGFIdentityInformation
	err := ngapcodec.DecodeParameter("W-AGFIdentityInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return WAGFIdentityInformation{}, 0, err
	}
	return v, len(data), err
}
