// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TimeSinceFailure represents ASN.1 TimeSinceFailure.
type TimeSinceFailure struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:172800"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TimeSinceFailure) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TimeSinceFailure", v)
}

// DecodeTimeSinceFailure decodes a standalone APER value and reports the consumed octets.
func DecodeTimeSinceFailure(data []byte) (TimeSinceFailure, int, error) {
	var v TimeSinceFailure
	err := ngapcodec.DecodeParameter("TimeSinceFailure", data, &v)
	if ngapcodec.IsFatal(err) {
		return TimeSinceFailure{}, 0, err
	}
	return v, len(data), err
}
