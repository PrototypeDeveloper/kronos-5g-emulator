// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FurtherExtendedUEIdentityIndexValue represents ASN.1 FurtherExtendedUEIdentityIndexValue.
type FurtherExtendedUEIdentityIndexValue struct {
	Value BitString `aper:"sizeLB:20,sizeUB:20"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FurtherExtendedUEIdentityIndexValue) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FurtherExtendedUEIdentityIndexValue", v)
}

// DecodeFurtherExtendedUEIdentityIndexValue decodes a standalone APER value and reports the consumed octets.
func DecodeFurtherExtendedUEIdentityIndexValue(data []byte) (FurtherExtendedUEIdentityIndexValue, int, error) {
	var v FurtherExtendedUEIdentityIndexValue
	err := ngapcodec.DecodeParameter("FurtherExtendedUEIdentityIndexValue", data, &v)
	if ngapcodec.IsFatal(err) {
		return FurtherExtendedUEIdentityIndexValue{}, 0, err
	}
	return v, len(data), err
}
