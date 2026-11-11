// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PagingPolicyDifferentiationList represents ASN.1 PagingPolicyDifferentiationList.
type PagingPolicyDifferentiationList struct {
	List []PagingPolicyDifferentiationItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PagingPolicyDifferentiationList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PagingPolicyDifferentiationList", v)
}

// DecodePagingPolicyDifferentiationList decodes a standalone APER value and reports the consumed octets.
func DecodePagingPolicyDifferentiationList(data []byte) (PagingPolicyDifferentiationList, int, error) {
	var v PagingPolicyDifferentiationList
	err := ngapcodec.DecodeParameter("PagingPolicyDifferentiationList", data, &v)
	if ngapcodec.IsFatal(err) {
		return PagingPolicyDifferentiationList{}, 0, err
	}
	return v, len(data), err
}
