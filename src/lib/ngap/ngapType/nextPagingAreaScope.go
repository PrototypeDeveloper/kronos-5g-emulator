// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NextPagingAreaScope represents ASN.1 NextPagingAreaScope.
const (
	NextPagingAreaScopePresentSame    Enumerated = 0
	NextPagingAreaScopePresentChanged Enumerated = 1
)

type NextPagingAreaScope struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NextPagingAreaScope) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NextPagingAreaScope", v)
}

// DecodeNextPagingAreaScope decodes a standalone APER value and reports the consumed octets.
func DecodeNextPagingAreaScope(data []byte) (NextPagingAreaScope, int, error) {
	var v NextPagingAreaScope
	err := ngapcodec.DecodeParameter("NextPagingAreaScope", data, &v)
	if ngapcodec.IsFatal(err) {
		return NextPagingAreaScope{}, 0, err
	}
	return v, len(data), err
}
