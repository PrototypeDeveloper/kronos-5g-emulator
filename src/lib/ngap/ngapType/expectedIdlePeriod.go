// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExpectedIdlePeriod represents ASN.1 ExpectedIdlePeriod.
type ExpectedIdlePeriod struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:181"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExpectedIdlePeriod) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExpectedIdlePeriod", v)
}

// DecodeExpectedIdlePeriod decodes a standalone APER value and reports the consumed octets.
func DecodeExpectedIdlePeriod(data []byte) (ExpectedIdlePeriod, int, error) {
	var v ExpectedIdlePeriod
	err := ngapcodec.DecodeParameter("ExpectedIdlePeriod", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExpectedIdlePeriod{}, 0, err
	}
	return v, len(data), err
}
