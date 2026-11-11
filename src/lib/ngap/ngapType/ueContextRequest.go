// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEContextRequest represents ASN.1 UEContextRequest.
const (
	UEContextRequestPresentRequested Enumerated = 0
)

type UEContextRequest struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEContextRequest) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEContextRequest", v)
}

// DecodeUEContextRequest decodes a standalone APER value and reports the consumed octets.
func DecodeUEContextRequest(data []byte) (UEContextRequest, int, error) {
	var v UEContextRequest
	err := ngapcodec.DecodeParameter("UEContextRequest", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEContextRequest{}, 0, err
	}
	return v, len(data), err
}
