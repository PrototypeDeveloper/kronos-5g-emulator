// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SCGActivationTime represents ASN.1 SCGActivationTime.
type SCGActivationTime struct {
	Value int64 `aper:"valueLB:0,valueUB:40950"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SCGActivationTime) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SCGActivationTime", v)
}

// DecodeSCGActivationTime decodes a standalone APER value and reports the consumed octets.
func DecodeSCGActivationTime(data []byte) (SCGActivationTime, int, error) {
	var v SCGActivationTime
	err := ngapcodec.DecodeParameter("SCGActivationTime", data, &v)
	if ngapcodec.IsFatal(err) {
		return SCGActivationTime{}, 0, err
	}
	return v, len(data), err
}
