// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DelayCritical represents ASN.1 DelayCritical.
const (
	DelayCriticalPresentDelayCritical    Enumerated = 0
	DelayCriticalPresentNonDelayCritical Enumerated = 1
)

type DelayCritical struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DelayCritical) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("DelayCritical", v) }

// DecodeDelayCritical decodes a standalone APER value and reports the consumed octets.
func DecodeDelayCritical(data []byte) (DelayCritical, int, error) {
	var v DelayCritical
	err := ngapcodec.DecodeParameter("DelayCritical", data, &v)
	if ngapcodec.IsFatal(err) {
		return DelayCritical{}, 0, err
	}
	return v, len(data), err
}
