// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NextHopChainingCount represents ASN.1 NextHopChainingCount.
type NextHopChainingCount struct {
	Value int64 `aper:"valueLB:0,valueUB:7"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NextHopChainingCount) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NextHopChainingCount", v)
}

// DecodeNextHopChainingCount decodes a standalone APER value and reports the consumed octets.
func DecodeNextHopChainingCount(data []byte) (NextHopChainingCount, int, error) {
	var v NextHopChainingCount
	err := ngapcodec.DecodeParameter("NextHopChainingCount", data, &v)
	if ngapcodec.IsFatal(err) {
		return NextHopChainingCount{}, 0, err
	}
	return v, len(data), err
}
