// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// Hysteresis represents ASN.1 Hysteresis.
type Hysteresis struct {
	Value int64 `aper:"valueLB:0,valueUB:30"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v Hysteresis) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("Hysteresis", v) }

// DecodeHysteresis decodes a standalone APER value and reports the consumed octets.
func DecodeHysteresis(data []byte) (Hysteresis, int, error) {
	var v Hysteresis
	err := ngapcodec.DecodeParameter("Hysteresis", data, &v)
	if ngapcodec.IsFatal(err) {
		return Hysteresis{}, 0, err
	}
	return v, len(data), err
}
