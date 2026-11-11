// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PagingDRX represents ASN.1 PagingDRX.
const (
	PagingDRXPresentV32  Enumerated = 0
	PagingDRXPresentV64  Enumerated = 1
	PagingDRXPresentV128 Enumerated = 2
	PagingDRXPresentV256 Enumerated = 3
)

type PagingDRX struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PagingDRX) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("PagingDRX", v) }

// DecodePagingDRX decodes a standalone APER value and reports the consumed octets.
func DecodePagingDRX(data []byte) (PagingDRX, int, error) {
	var v PagingDRX
	err := ngapcodec.DecodeParameter("PagingDRX", data, &v)
	if ngapcodec.IsFatal(err) {
		return PagingDRX{}, 0, err
	}
	return v, len(data), err
}
