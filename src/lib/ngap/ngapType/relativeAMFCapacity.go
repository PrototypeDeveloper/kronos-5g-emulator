// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RelativeAMFCapacity represents ASN.1 RelativeAMFCapacity.
type RelativeAMFCapacity struct {
	Value int64 `aper:"valueLB:0,valueUB:255"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RelativeAMFCapacity) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RelativeAMFCapacity", v)
}

// DecodeRelativeAMFCapacity decodes a standalone APER value and reports the consumed octets.
func DecodeRelativeAMFCapacity(data []byte) (RelativeAMFCapacity, int, error) {
	var v RelativeAMFCapacity
	err := ngapcodec.DecodeParameter("RelativeAMFCapacity", data, &v)
	if ngapcodec.IsFatal(err) {
		return RelativeAMFCapacity{}, 0, err
	}
	return v, len(data), err
}
