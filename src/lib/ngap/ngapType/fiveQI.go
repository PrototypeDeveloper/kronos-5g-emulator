// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveQI represents ASN.1 FiveQI.
type FiveQI struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:255"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveQI) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("FiveQI", v) }

// DecodeFiveQI decodes a standalone APER value and reports the consumed octets.
func DecodeFiveQI(data []byte) (FiveQI, int, error) {
	var v FiveQI
	err := ngapcodec.DecodeParameter("FiveQI", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveQI{}, 0, err
	}
	return v, len(data), err
}
