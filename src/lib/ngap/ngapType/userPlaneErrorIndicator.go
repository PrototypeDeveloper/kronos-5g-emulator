// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UserPlaneErrorIndicator represents ASN.1 UserPlaneErrorIndicator.
const (
	UserPlaneErrorIndicatorPresentGTPUErrorIndicationReceived Enumerated = 0
)

type UserPlaneErrorIndicator struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserPlaneErrorIndicator) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserPlaneErrorIndicator", v)
}

// DecodeUserPlaneErrorIndicator decodes a standalone APER value and reports the consumed octets.
func DecodeUserPlaneErrorIndicator(data []byte) (UserPlaneErrorIndicator, int, error) {
	var v UserPlaneErrorIndicator
	err := ngapcodec.DecodeParameter("UserPlaneErrorIndicator", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserPlaneErrorIndicator{}, 0, err
	}
	return v, len(data), err
}
