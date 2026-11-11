// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ECNMarkingAtUPFRequest represents ASN.1 ECNMarkingAtUPFRequest.
const (
	ECNMarkingAtUPFRequestPresentUl   Enumerated = 0
	ECNMarkingAtUPFRequestPresentDl   Enumerated = 1
	ECNMarkingAtUPFRequestPresentBoth Enumerated = 2
	ECNMarkingAtUPFRequestPresentStop Enumerated = 3
)

type ECNMarkingAtUPFRequest struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ECNMarkingAtUPFRequest) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ECNMarkingAtUPFRequest", v)
}

// DecodeECNMarkingAtUPFRequest decodes a standalone APER value and reports the consumed octets.
func DecodeECNMarkingAtUPFRequest(data []byte) (ECNMarkingAtUPFRequest, int, error) {
	var v ECNMarkingAtUPFRequest
	err := ngapcodec.DecodeParameter("ECNMarkingAtUPFRequest", data, &v)
	if ngapcodec.IsFatal(err) {
		return ECNMarkingAtUPFRequest{}, 0, err
	}
	return v, len(data), err
}
