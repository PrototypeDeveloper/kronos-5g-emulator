// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SLPositioningRangingAuthorized represents ASN.1 SLPositioningRangingAuthorized.
const (
	SLPositioningRangingAuthorizedPresentAuthorized    Enumerated = 0
	SLPositioningRangingAuthorizedPresentNotAuthorized Enumerated = 1
)

type SLPositioningRangingAuthorized struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SLPositioningRangingAuthorized) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SLPositioningRangingAuthorized", v)
}

// DecodeSLPositioningRangingAuthorized decodes a standalone APER value and reports the consumed octets.
func DecodeSLPositioningRangingAuthorized(data []byte) (SLPositioningRangingAuthorized, int, error) {
	var v SLPositioningRangingAuthorized
	err := ngapcodec.DecodeParameter("SLPositioningRangingAuthorized", data, &v)
	if ngapcodec.IsFatal(err) {
		return SLPositioningRangingAuthorized{}, 0, err
	}
	return v, len(data), err
}
