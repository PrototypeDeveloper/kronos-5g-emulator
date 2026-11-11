// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TimeToWait represents ASN.1 TimeToWait.
const (
	TimeToWaitPresentV1s  Enumerated = 0
	TimeToWaitPresentV2s  Enumerated = 1
	TimeToWaitPresentV5s  Enumerated = 2
	TimeToWaitPresentV10s Enumerated = 3
	TimeToWaitPresentV20s Enumerated = 4
	TimeToWaitPresentV60s Enumerated = 5
)

type TimeToWait struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:5"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TimeToWait) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TimeToWait", v) }

// DecodeTimeToWait decodes a standalone APER value and reports the consumed octets.
func DecodeTimeToWait(data []byte) (TimeToWait, int, error) {
	var v TimeToWait
	err := ngapcodec.DecodeParameter("TimeToWait", data, &v)
	if ngapcodec.IsFatal(err) {
		return TimeToWait{}, 0, err
	}
	return v, len(data), err
}
