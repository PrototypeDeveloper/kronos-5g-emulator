// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PagingAttemptCount represents ASN.1 PagingAttemptCount.
type PagingAttemptCount struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PagingAttemptCount) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PagingAttemptCount", v)
}

// DecodePagingAttemptCount decodes a standalone APER value and reports the consumed octets.
func DecodePagingAttemptCount(data []byte) (PagingAttemptCount, int, error) {
	var v PagingAttemptCount
	err := ngapcodec.DecodeParameter("PagingAttemptCount", data, &v)
	if ngapcodec.IsFatal(err) {
		return PagingAttemptCount{}, 0, err
	}
	return v, len(data), err
}
