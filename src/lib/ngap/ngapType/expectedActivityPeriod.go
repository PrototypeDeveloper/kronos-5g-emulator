// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExpectedActivityPeriod represents ASN.1 ExpectedActivityPeriod.
type ExpectedActivityPeriod struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:181"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExpectedActivityPeriod) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExpectedActivityPeriod", v)
}

// DecodeExpectedActivityPeriod decodes a standalone APER value and reports the consumed octets.
func DecodeExpectedActivityPeriod(data []byte) (ExpectedActivityPeriod, int, error) {
	var v ExpectedActivityPeriod
	err := ngapcodec.DecodeParameter("ExpectedActivityPeriod", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExpectedActivityPeriod{}, 0, err
	}
	return v, len(data), err
}
