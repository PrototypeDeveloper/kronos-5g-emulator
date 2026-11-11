// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ForbiddenAreaInformation represents ASN.1 ForbiddenAreaInformation.
type ForbiddenAreaInformation struct {
	List []ForbiddenAreaInformationItem `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ForbiddenAreaInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ForbiddenAreaInformation", v)
}

// DecodeForbiddenAreaInformation decodes a standalone APER value and reports the consumed octets.
func DecodeForbiddenAreaInformation(data []byte) (ForbiddenAreaInformation, int, error) {
	var v ForbiddenAreaInformation
	err := ngapcodec.DecodeParameter("ForbiddenAreaInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return ForbiddenAreaInformation{}, 0, err
	}
	return v, len(data), err
}
