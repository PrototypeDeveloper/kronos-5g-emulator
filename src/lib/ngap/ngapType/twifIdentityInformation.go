// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TWIFIdentityInformation represents ASN.1 TWIFIdentityInformation.
type TWIFIdentityInformation struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TWIFIdentityInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TWIFIdentityInformation", v)
}

// DecodeTWIFIdentityInformation decodes a standalone APER value and reports the consumed octets.
func DecodeTWIFIdentityInformation(data []byte) (TWIFIdentityInformation, int, error) {
	var v TWIFIdentityInformation
	err := ngapcodec.DecodeParameter("TWIFIdentityInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return TWIFIdentityInformation{}, 0, err
	}
	return v, len(data), err
}
