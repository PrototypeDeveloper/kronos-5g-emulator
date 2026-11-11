// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ECNMarkingAtRANRequest represents ASN.1 ECNMarkingAtRANRequest.
const (
	ECNMarkingAtRANRequestPresentUl   Enumerated = 0
	ECNMarkingAtRANRequestPresentDl   Enumerated = 1
	ECNMarkingAtRANRequestPresentBoth Enumerated = 2
	ECNMarkingAtRANRequestPresentStop Enumerated = 3
)

type ECNMarkingAtRANRequest struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ECNMarkingAtRANRequest) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ECNMarkingAtRANRequest", v)
}

// DecodeECNMarkingAtRANRequest decodes a standalone APER value and reports the consumed octets.
func DecodeECNMarkingAtRANRequest(data []byte) (ECNMarkingAtRANRequest, int, error) {
	var v ECNMarkingAtRANRequest
	err := ngapcodec.DecodeParameter("ECNMarkingAtRANRequest", data, &v)
	if ngapcodec.IsFatal(err) {
		return ECNMarkingAtRANRequest{}, 0, err
	}
	return v, len(data), err
}
