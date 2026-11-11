// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ThresholdSINR represents ASN.1 Threshold-SINR.
type ThresholdSINR struct {
	Value int64 `aper:"valueLB:0,valueUB:127"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ThresholdSINR) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Threshold-SINR", v)
}

// DecodeThresholdSINR decodes a standalone APER value and reports the consumed octets.
func DecodeThresholdSINR(data []byte) (ThresholdSINR, int, error) {
	var v ThresholdSINR
	err := ngapcodec.DecodeParameter("Threshold-SINR", data, &v)
	if ngapcodec.IsFatal(err) {
		return ThresholdSINR{}, 0, err
	}
	return v, len(data), err
}
