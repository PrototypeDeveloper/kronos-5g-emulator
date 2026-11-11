// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RepetitionPeriod represents ASN.1 RepetitionPeriod.
type RepetitionPeriod struct {
	Value int64 `aper:"valueLB:0,valueUB:131071"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RepetitionPeriod) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RepetitionPeriod", v)
}

// DecodeRepetitionPeriod decodes a standalone APER value and reports the consumed octets.
func DecodeRepetitionPeriod(data []byte) (RepetitionPeriod, int, error) {
	var v RepetitionPeriod
	err := ngapcodec.DecodeParameter("RepetitionPeriod", data, &v)
	if ngapcodec.IsFatal(err) {
		return RepetitionPeriod{}, 0, err
	}
	return v, len(data), err
}
