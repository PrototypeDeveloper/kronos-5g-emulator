// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NumberOfBroadcastsRequested represents ASN.1 NumberOfBroadcastsRequested.
type NumberOfBroadcastsRequested struct {
	Value int64 `aper:"valueLB:0,valueUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NumberOfBroadcastsRequested) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NumberOfBroadcastsRequested", v)
}

// DecodeNumberOfBroadcastsRequested decodes a standalone APER value and reports the consumed octets.
func DecodeNumberOfBroadcastsRequested(data []byte) (NumberOfBroadcastsRequested, int, error) {
	var v NumberOfBroadcastsRequested
	err := ngapcodec.DecodeParameter("NumberOfBroadcastsRequested", data, &v)
	if ngapcodec.IsFatal(err) {
		return NumberOfBroadcastsRequested{}, 0, err
	}
	return v, len(data), err
}
