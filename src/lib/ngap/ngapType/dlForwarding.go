// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DLForwarding represents ASN.1 DLForwarding.
const (
	DLForwardingPresentDlForwardingProposed Enumerated = 0
)

type DLForwarding struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DLForwarding) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("DLForwarding", v) }

// DecodeDLForwarding decodes a standalone APER value and reports the consumed octets.
func DecodeDLForwarding(data []byte) (DLForwarding, int, error) {
	var v DLForwarding
	err := ngapcodec.DecodeParameter("DLForwarding", data, &v)
	if ngapcodec.IsFatal(err) {
		return DLForwarding{}, 0, err
	}
	return v, len(data), err
}
