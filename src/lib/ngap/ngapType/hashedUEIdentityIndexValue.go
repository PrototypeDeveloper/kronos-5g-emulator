// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// HashedUEIdentityIndexValue represents ASN.1 HashedUEIdentityIndexValue.
type HashedUEIdentityIndexValue struct {
	Value BitString `aper:"sizeExt,sizeLB:13,sizeUB:13"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HashedUEIdentityIndexValue) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("HashedUEIdentityIndexValue", v)
}

// DecodeHashedUEIdentityIndexValue decodes a standalone APER value and reports the consumed octets.
func DecodeHashedUEIdentityIndexValue(data []byte) (HashedUEIdentityIndexValue, int, error) {
	var v HashedUEIdentityIndexValue
	err := ngapcodec.DecodeParameter("HashedUEIdentityIndexValue", data, &v)
	if ngapcodec.IsFatal(err) {
		return HashedUEIdentityIndexValue{}, 0, err
	}
	return v, len(data), err
}
