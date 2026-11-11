// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SONInformationRequest represents ASN.1 SONInformationRequest.
const (
	SONInformationRequestPresentXnTNLConfigurationInfo Enumerated = 0
)

type SONInformationRequest struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SONInformationRequest) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SONInformationRequest", v)
}

// DecodeSONInformationRequest decodes a standalone APER value and reports the consumed octets.
func DecodeSONInformationRequest(data []byte) (SONInformationRequest, int, error) {
	var v SONInformationRequest
	err := ngapcodec.DecodeParameter("SONInformationRequest", data, &v)
	if ngapcodec.IsFatal(err) {
		return SONInformationRequest{}, 0, err
	}
	return v, len(data), err
}
