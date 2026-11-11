// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MobilityInformation represents ASN.1 MobilityInformation.
type MobilityInformation struct {
	Value BitString `aper:"sizeLB:16,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MobilityInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MobilityInformation", v)
}

// DecodeMobilityInformation decodes a standalone APER value and reports the consumed octets.
func DecodeMobilityInformation(data []byte) (MobilityInformation, int, error) {
	var v MobilityInformation
	err := ngapcodec.DecodeParameter("MobilityInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return MobilityInformation{}, 0, err
	}
	return v, len(data), err
}
