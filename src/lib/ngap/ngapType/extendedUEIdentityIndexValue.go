// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExtendedUEIdentityIndexValue represents ASN.1 ExtendedUEIdentityIndexValue.
type ExtendedUEIdentityIndexValue struct {
	Value BitString `aper:"sizeLB:16,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExtendedUEIdentityIndexValue) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExtendedUEIdentityIndexValue", v)
}

// DecodeExtendedUEIdentityIndexValue decodes a standalone APER value and reports the consumed octets.
func DecodeExtendedUEIdentityIndexValue(data []byte) (ExtendedUEIdentityIndexValue, int, error) {
	var v ExtendedUEIdentityIndexValue
	err := ngapcodec.DecodeParameter("ExtendedUEIdentityIndexValue", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExtendedUEIdentityIndexValue{}, 0, err
	}
	return v, len(data), err
}
