// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EUTRACellIdentity represents ASN.1 EUTRACellIdentity.
type EUTRACellIdentity struct {
	Value BitString `aper:"sizeLB:28,sizeUB:28"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRACellIdentity) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRACellIdentity", v)
}

// DecodeEUTRACellIdentity decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRACellIdentity(data []byte) (EUTRACellIdentity, int, error) {
	var v EUTRACellIdentity
	err := ngapcodec.DecodeParameter("EUTRACellIdentity", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRACellIdentity{}, 0, err
	}
	return v, len(data), err
}
