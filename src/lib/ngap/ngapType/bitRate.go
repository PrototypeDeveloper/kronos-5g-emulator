// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BitRate represents ASN.1 BitRate.
type BitRate struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:4000000000000"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BitRate) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("BitRate", v) }

// DecodeBitRate decodes a standalone APER value and reports the consumed octets.
func DecodeBitRate(data []byte) (BitRate, int, error) {
	var v BitRate
	err := ngapcodec.DecodeParameter("BitRate", data, &v)
	if ngapcodec.IsFatal(err) {
		return BitRate{}, 0, err
	}
	return v, len(data), err
}
