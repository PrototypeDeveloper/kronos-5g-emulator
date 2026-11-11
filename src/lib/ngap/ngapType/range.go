// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// Range represents ASN.1 Range.
const (
	RangePresentM50   Enumerated = 0
	RangePresentM80   Enumerated = 1
	RangePresentM180  Enumerated = 2
	RangePresentM200  Enumerated = 3
	RangePresentM350  Enumerated = 4
	RangePresentM400  Enumerated = 5
	RangePresentM500  Enumerated = 6
	RangePresentM700  Enumerated = 7
	RangePresentM1000 Enumerated = 8
)

type Range struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v Range) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("Range", v) }

// DecodeRange decodes a standalone APER value and reports the consumed octets.
func DecodeRange(data []byte) (Range, int, error) {
	var v Range
	err := ngapcodec.DecodeParameter("Range", data, &v)
	if ngapcodec.IsFatal(err) {
		return Range{}, 0, err
	}
	return v, len(data), err
}
