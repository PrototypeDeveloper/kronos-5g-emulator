// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CongestionInformationRequest represents ASN.1 CongestionInformationRequest.
const (
	CongestionInformationRequestPresentUl   Enumerated = 0
	CongestionInformationRequestPresentDl   Enumerated = 1
	CongestionInformationRequestPresentBoth Enumerated = 2
	CongestionInformationRequestPresentStop Enumerated = 3
)

type CongestionInformationRequest struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CongestionInformationRequest) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CongestionInformationRequest", v)
}

// DecodeCongestionInformationRequest decodes a standalone APER value and reports the consumed octets.
func DecodeCongestionInformationRequest(data []byte) (CongestionInformationRequest, int, error) {
	var v CongestionInformationRequest
	err := ngapcodec.DecodeParameter("CongestionInformationRequest", data, &v)
	if ngapcodec.IsFatal(err) {
		return CongestionInformationRequest{}, 0, err
	}
	return v, len(data), err
}
