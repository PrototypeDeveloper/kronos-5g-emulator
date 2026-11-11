// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NASSecurityParametersFromNGRAN represents ASN.1 NASSecurityParametersFromNGRAN.
type NASSecurityParametersFromNGRAN struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NASSecurityParametersFromNGRAN) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NASSecurityParametersFromNGRAN", v)
}

// DecodeNASSecurityParametersFromNGRAN decodes a standalone APER value and reports the consumed octets.
func DecodeNASSecurityParametersFromNGRAN(data []byte) (NASSecurityParametersFromNGRAN, int, error) {
	var v NASSecurityParametersFromNGRAN
	err := ngapcodec.DecodeParameter("NASSecurityParametersFromNGRAN", data, &v)
	if ngapcodec.IsFatal(err) {
		return NASSecurityParametersFromNGRAN{}, 0, err
	}
	return v, len(data), err
}
