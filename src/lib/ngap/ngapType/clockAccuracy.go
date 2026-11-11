// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ClockAccuracy represents ASN.1 ClockAccuracy.
type ClockAccuracyPresent uint32

const (
	ClockAccuracyPresentNothing ClockAccuracyPresent = iota
	ClockAccuracyPresentClockAccuracyValue
	ClockAccuracyPresentClockAccuracyIndex
	ClockAccuracyPresentChoiceExtensions
)

type ClockAccuracy struct {
	Present            ClockAccuracyPresent
	ClockAccuracyValue int64 `aper:"valueExt,valueLB:1,valueUB:40000000"`
	ClockAccuracyIndex int64 `aper:"valueExt,valueLB:32,valueUB:47"`
	ChoiceExtensions   ClockAccuracyExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ClockAccuracy) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("ClockAccuracy", v) }

// DecodeClockAccuracy decodes a standalone APER value and reports the consumed octets.
func DecodeClockAccuracy(data []byte) (ClockAccuracy, int, error) {
	var v ClockAccuracy
	err := ngapcodec.DecodeParameter("ClockAccuracy", data, &v)
	if ngapcodec.IsFatal(err) {
		return ClockAccuracy{}, 0, err
	}
	return v, len(data), err
}
