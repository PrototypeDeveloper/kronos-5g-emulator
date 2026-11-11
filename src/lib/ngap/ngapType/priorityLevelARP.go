// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PriorityLevelARP represents ASN.1 PriorityLevelARP.
type PriorityLevelARP struct {
	Value int64 `aper:"valueLB:1,valueUB:15"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PriorityLevelARP) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PriorityLevelARP", v)
}

// DecodePriorityLevelARP decodes a standalone APER value and reports the consumed octets.
func DecodePriorityLevelARP(data []byte) (PriorityLevelARP, int, error) {
	var v PriorityLevelARP
	err := ngapcodec.DecodeParameter("PriorityLevelARP", data, &v)
	if ngapcodec.IsFatal(err) {
		return PriorityLevelARP{}, 0, err
	}
	return v, len(data), err
}
