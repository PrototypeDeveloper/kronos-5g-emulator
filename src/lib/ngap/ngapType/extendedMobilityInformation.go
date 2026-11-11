// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExtendedMobilityInformation represents ASN.1 ExtendedMobilityInformation.
type ExtendedMobilityInformation struct {
	Value BitString `aper:"sizeLB:32,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExtendedMobilityInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExtendedMobilityInformation", v)
}

// DecodeExtendedMobilityInformation decodes a standalone APER value and reports the consumed octets.
func DecodeExtendedMobilityInformation(data []byte) (ExtendedMobilityInformation, int, error) {
	var v ExtendedMobilityInformation
	err := ngapcodec.DecodeParameter("ExtendedMobilityInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExtendedMobilityInformation{}, 0, err
	}
	return v, len(data), err
}
