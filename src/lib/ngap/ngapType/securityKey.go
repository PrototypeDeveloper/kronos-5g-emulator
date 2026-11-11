// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SecurityKey represents ASN.1 SecurityKey.
type SecurityKey struct {
	Value BitString `aper:"sizeLB:256,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SecurityKey) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("SecurityKey", v) }

// DecodeSecurityKey decodes a standalone APER value and reports the consumed octets.
func DecodeSecurityKey(data []byte) (SecurityKey, int, error) {
	var v SecurityKey
	err := ngapcodec.DecodeParameter("SecurityKey", data, &v)
	if ngapcodec.IsFatal(err) {
		return SecurityKey{}, 0, err
	}
	return v, len(data), err
}
