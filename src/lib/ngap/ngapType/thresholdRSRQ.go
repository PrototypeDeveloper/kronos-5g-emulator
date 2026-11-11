// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ThresholdRSRQ represents ASN.1 Threshold-RSRQ.
type ThresholdRSRQ struct {
	Value int64 `aper:"valueLB:0,valueUB:127"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ThresholdRSRQ) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Threshold-RSRQ", v)
}

// DecodeThresholdRSRQ decodes a standalone APER value and reports the consumed octets.
func DecodeThresholdRSRQ(data []byte) (ThresholdRSRQ, int, error) {
	var v ThresholdRSRQ
	err := ngapcodec.DecodeParameter("Threshold-RSRQ", data, &v)
	if ngapcodec.IsFatal(err) {
		return ThresholdRSRQ{}, 0, err
	}
	return v, len(data), err
}
