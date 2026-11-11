// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// M7period represents ASN.1 M7period.
type M7period struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:60"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M7period) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("M7period", v) }

// DecodeM7period decodes a standalone APER value and reports the consumed octets.
func DecodeM7period(data []byte) (M7period, int, error) {
	var v M7period
	err := ngapcodec.DecodeParameter("M7period", data, &v)
	if ngapcodec.IsFatal(err) {
		return M7period{}, 0, err
	}
	return v, len(data), err
}
