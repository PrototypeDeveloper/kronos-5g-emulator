// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ULForwarding represents ASN.1 ULForwarding.
const (
	ULForwardingPresentUlForwardingProposed Enumerated = 0
)

type ULForwarding struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ULForwarding) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("ULForwarding", v) }

// DecodeULForwarding decodes a standalone APER value and reports the consumed octets.
func DecodeULForwarding(data []byte) (ULForwarding, int, error) {
	var v ULForwarding
	err := ngapcodec.DecodeParameter("ULForwarding", data, &v)
	if ngapcodec.IsFatal(err) {
		return ULForwarding{}, 0, err
	}
	return v, len(data), err
}
