// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// HandoverWindowStart represents ASN.1 HandoverWindowStart.
type HandoverWindowStart struct {
	Value int64 `aper:"valueLB:0,valueUB:549755813887"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HandoverWindowStart) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("HandoverWindowStart", v)
}

// DecodeHandoverWindowStart decodes a standalone APER value and reports the consumed octets.
func DecodeHandoverWindowStart(data []byte) (HandoverWindowStart, int, error) {
	var v HandoverWindowStart
	err := ngapcodec.DecodeParameter("HandoverWindowStart", data, &v)
	if ngapcodec.IsFatal(err) {
		return HandoverWindowStart{}, 0, err
	}
	return v, len(data), err
}
