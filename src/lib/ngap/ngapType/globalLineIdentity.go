// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GlobalLineIdentity represents ASN.1 GlobalLineIdentity.
type GlobalLineIdentity struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalLineIdentity) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GlobalLineIdentity", v)
}

// DecodeGlobalLineIdentity decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalLineIdentity(data []byte) (GlobalLineIdentity, int, error) {
	var v GlobalLineIdentity
	err := ngapcodec.DecodeParameter("GlobalLineIdentity", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalLineIdentity{}, 0, err
	}
	return v, len(data), err
}
