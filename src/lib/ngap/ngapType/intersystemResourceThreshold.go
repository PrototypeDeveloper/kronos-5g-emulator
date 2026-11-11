// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntersystemResourceThreshold represents ASN.1 IntersystemResourceThreshold.
type IntersystemResourceThreshold struct {
	Value int64 `aper:"valueLB:0,valueUB:100"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemResourceThreshold) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemResourceThreshold", v)
}

// DecodeIntersystemResourceThreshold decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemResourceThreshold(data []byte) (IntersystemResourceThreshold, int, error) {
	var v IntersystemResourceThreshold
	err := ngapcodec.DecodeParameter("IntersystemResourceThreshold", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemResourceThreshold{}, 0, err
	}
	return v, len(data), err
}
