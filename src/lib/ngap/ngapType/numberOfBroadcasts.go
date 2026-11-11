// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NumberOfBroadcasts represents ASN.1 NumberOfBroadcasts.
type NumberOfBroadcasts struct {
	Value int64 `aper:"valueLB:0,valueUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NumberOfBroadcasts) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NumberOfBroadcasts", v)
}

// DecodeNumberOfBroadcasts decodes a standalone APER value and reports the consumed octets.
func DecodeNumberOfBroadcasts(data []byte) (NumberOfBroadcasts, int, error) {
	var v NumberOfBroadcasts
	err := ngapcodec.DecodeParameter("NumberOfBroadcasts", data, &v)
	if ngapcodec.IsFatal(err) {
		return NumberOfBroadcasts{}, 0, err
	}
	return v, len(data), err
}
