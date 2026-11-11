// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// M6reportInterval represents ASN.1 M6report-Interval.
const (
	M6reportIntervalPresentMs120   Enumerated = 0
	M6reportIntervalPresentMs240   Enumerated = 1
	M6reportIntervalPresentMs480   Enumerated = 2
	M6reportIntervalPresentMs640   Enumerated = 3
	M6reportIntervalPresentMs1024  Enumerated = 4
	M6reportIntervalPresentMs2048  Enumerated = 5
	M6reportIntervalPresentMs5120  Enumerated = 6
	M6reportIntervalPresentMs10240 Enumerated = 7
	M6reportIntervalPresentMs20480 Enumerated = 8
	M6reportIntervalPresentMs40960 Enumerated = 9
	M6reportIntervalPresentMin1    Enumerated = 10
	M6reportIntervalPresentMin6    Enumerated = 11
	M6reportIntervalPresentMin12   Enumerated = 12
	M6reportIntervalPresentMin30   Enumerated = 13
)

type M6reportInterval struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:13"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M6reportInterval) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M6report-Interval", v)
}

// DecodeM6reportInterval decodes a standalone APER value and reports the consumed octets.
func DecodeM6reportInterval(data []byte) (M6reportInterval, int, error) {
	var v M6reportInterval
	err := ngapcodec.DecodeParameter("M6report-Interval", data, &v)
	if ngapcodec.IsFatal(err) {
		return M6reportInterval{}, 0, err
	}
	return v, len(data), err
}
