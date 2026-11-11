// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PagingPolicyDifferentiation represents ASN.1 PagingPolicyDifferentiation.
type PagingPolicyDifferentiation struct {
	PDUSessionForPagingList PDUSessionForPagingList
	HasIEExtensions         bool
	IEExtensions            []PagingPolicyDifferentiationExtIEsEntry
	UnknownExtensions       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PagingPolicyDifferentiation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PagingPolicyDifferentiation", v)
}

// DecodePagingPolicyDifferentiation decodes a standalone APER value and reports the consumed octets.
func DecodePagingPolicyDifferentiation(data []byte) (PagingPolicyDifferentiation, int, error) {
	var v PagingPolicyDifferentiation
	err := ngapcodec.DecodeParameter("PagingPolicyDifferentiation", data, &v)
	if ngapcodec.IsFatal(err) {
		return PagingPolicyDifferentiation{}, 0, err
	}
	return v, len(data), err
}
