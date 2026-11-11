// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ThresholdRSRP represents ASN.1 Threshold-RSRP.
type ThresholdRSRP struct {
	Value int64 `aper:"valueLB:0,valueUB:127"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ThresholdRSRP) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Threshold-RSRP", v)
}

// DecodeThresholdRSRP decodes a standalone APER value and reports the consumed octets.
func DecodeThresholdRSRP(data []byte) (ThresholdRSRP, int, error) {
	var v ThresholdRSRP
	err := ngapcodec.DecodeParameter("Threshold-RSRP", data, &v)
	if ngapcodec.IsFatal(err) {
		return ThresholdRSRP{}, 0, err
	}
	return v, len(data), err
}
