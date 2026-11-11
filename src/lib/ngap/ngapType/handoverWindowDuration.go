// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// HandoverWindowDuration represents ASN.1 HandoverWindowDuration.
type HandoverWindowDuration struct {
	Value int64 `aper:"valueLB:1,valueUB:6000"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HandoverWindowDuration) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("HandoverWindowDuration", v)
}

// DecodeHandoverWindowDuration decodes a standalone APER value and reports the consumed octets.
func DecodeHandoverWindowDuration(data []byte) (HandoverWindowDuration, int, error) {
	var v HandoverWindowDuration
	err := ngapcodec.DecodeParameter("HandoverWindowDuration", data, &v)
	if ngapcodec.IsFatal(err) {
		return HandoverWindowDuration{}, 0, err
	}
	return v, len(data), err
}
