// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SurvivalTime represents ASN.1 SurvivalTime.
type SurvivalTime struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:1920000"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SurvivalTime) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("SurvivalTime", v) }

// DecodeSurvivalTime decodes a standalone APER value and reports the consumed octets.
func DecodeSurvivalTime(data []byte) (SurvivalTime, int, error) {
	var v SurvivalTime
	err := ngapcodec.DecodeParameter("SurvivalTime", data, &v)
	if ngapcodec.IsFatal(err) {
		return SurvivalTime{}, 0, err
	}
	return v, len(data), err
}
