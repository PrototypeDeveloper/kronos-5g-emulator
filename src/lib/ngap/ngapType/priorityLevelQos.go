// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PriorityLevelQos represents ASN.1 PriorityLevelQos.
type PriorityLevelQos struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:127"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PriorityLevelQos) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PriorityLevelQos", v)
}

// DecodePriorityLevelQos decodes a standalone APER value and reports the consumed octets.
func DecodePriorityLevelQos(data []byte) (PriorityLevelQos, int, error) {
	var v PriorityLevelQos
	err := ngapcodec.DecodeParameter("PriorityLevelQos", data, &v)
	if ngapcodec.IsFatal(err) {
		return PriorityLevelQos{}, 0, err
	}
	return v, len(data), err
}
