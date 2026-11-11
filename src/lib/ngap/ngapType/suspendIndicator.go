// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SuspendIndicator represents ASN.1 SuspendIndicator.
const (
	SuspendIndicatorPresentTrue Enumerated = 0
)

type SuspendIndicator struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SuspendIndicator) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SuspendIndicator", v)
}

// DecodeSuspendIndicator decodes a standalone APER value and reports the consumed octets.
func DecodeSuspendIndicator(data []byte) (SuspendIndicator, int, error) {
	var v SuspendIndicator
	err := ngapcodec.DecodeParameter("SuspendIndicator", data, &v)
	if ngapcodec.IsFatal(err) {
		return SuspendIndicator{}, 0, err
	}
	return v, len(data), err
}
