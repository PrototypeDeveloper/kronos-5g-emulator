// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PagingCause represents ASN.1 PagingCause.
const (
	PagingCausePresentVoice Enumerated = 0
)

type PagingCause struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PagingCause) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("PagingCause", v) }

// DecodePagingCause decodes a standalone APER value and reports the consumed octets.
func DecodePagingCause(data []byte) (PagingCause, int, error) {
	var v PagingCause
	err := ngapcodec.DecodeParameter("PagingCause", data, &v)
	if ngapcodec.IsFatal(err) {
		return PagingCause{}, 0, err
	}
	return v, len(data), err
}
