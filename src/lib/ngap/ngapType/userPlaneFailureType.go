// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UserPlaneFailureType represents ASN.1 UserPlaneFailureType.
const (
	UserPlaneFailureTypePresentGtpUErrorIndicationReceived Enumerated = 0
	UserPlaneFailureTypePresentUpPathFailure               Enumerated = 1
)

type UserPlaneFailureType struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserPlaneFailureType) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserPlaneFailureType", v)
}

// DecodeUserPlaneFailureType decodes a standalone APER value and reports the consumed octets.
func DecodeUserPlaneFailureType(data []byte) (UserPlaneFailureType, int, error) {
	var v UserPlaneFailureType
	err := ngapcodec.DecodeParameter("UserPlaneFailureType", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserPlaneFailureType{}, 0, err
	}
	return v, len(data), err
}
