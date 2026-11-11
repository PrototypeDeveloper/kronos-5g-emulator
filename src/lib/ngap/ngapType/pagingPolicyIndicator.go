// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PagingPolicyIndicator represents ASN.1 PagingPolicyIndicator.
type PagingPolicyIndicator struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:7"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PagingPolicyIndicator) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PagingPolicyIndicator", v)
}

// DecodePagingPolicyIndicator decodes a standalone APER value and reports the consumed octets.
func DecodePagingPolicyIndicator(data []byte) (PagingPolicyIndicator, int, error) {
	var v PagingPolicyIndicator
	err := ngapcodec.DecodeParameter("PagingPolicyIndicator", data, &v)
	if ngapcodec.IsFatal(err) {
		return PagingPolicyIndicator{}, 0, err
	}
	return v, len(data), err
}
