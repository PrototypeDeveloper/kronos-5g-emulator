// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PagingOrigin represents ASN.1 PagingOrigin.
const (
	PagingOriginPresentNon3gpp Enumerated = 0
)

type PagingOrigin struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PagingOrigin) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("PagingOrigin", v) }

// DecodePagingOrigin decodes a standalone APER value and reports the consumed octets.
func DecodePagingOrigin(data []byte) (PagingOrigin, int, error) {
	var v PagingOrigin
	err := ngapcodec.DecodeParameter("PagingOrigin", data, &v)
	if ngapcodec.IsFatal(err) {
		return PagingOrigin{}, 0, err
	}
	return v, len(data), err
}
