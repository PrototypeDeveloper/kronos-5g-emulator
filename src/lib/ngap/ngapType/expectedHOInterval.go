// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExpectedHOInterval represents ASN.1 ExpectedHOInterval.
const (
	ExpectedHOIntervalPresentSec15    Enumerated = 0
	ExpectedHOIntervalPresentSec30    Enumerated = 1
	ExpectedHOIntervalPresentSec60    Enumerated = 2
	ExpectedHOIntervalPresentSec90    Enumerated = 3
	ExpectedHOIntervalPresentSec120   Enumerated = 4
	ExpectedHOIntervalPresentSec180   Enumerated = 5
	ExpectedHOIntervalPresentLongTime Enumerated = 6
)

type ExpectedHOInterval struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:6"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExpectedHOInterval) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExpectedHOInterval", v)
}

// DecodeExpectedHOInterval decodes a standalone APER value and reports the consumed octets.
func DecodeExpectedHOInterval(data []byte) (ExpectedHOInterval, int, error) {
	var v ExpectedHOInterval
	err := ngapcodec.DecodeParameter("ExpectedHOInterval", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExpectedHOInterval{}, 0, err
	}
	return v, len(data), err
}
