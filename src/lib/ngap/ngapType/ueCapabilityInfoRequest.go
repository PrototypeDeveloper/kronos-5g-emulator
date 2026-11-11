// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UECapabilityInfoRequest represents ASN.1 UECapabilityInfoRequest.
const (
	UECapabilityInfoRequestPresentRequested Enumerated = 0
)

type UECapabilityInfoRequest struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UECapabilityInfoRequest) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UECapabilityInfoRequest", v)
}

// DecodeUECapabilityInfoRequest decodes a standalone APER value and reports the consumed octets.
func DecodeUECapabilityInfoRequest(data []byte) (UECapabilityInfoRequest, int, error) {
	var v UECapabilityInfoRequest
	err := ngapcodec.DecodeParameter("UECapabilityInfoRequest", data, &v)
	if ngapcodec.IsFatal(err) {
		return UECapabilityInfoRequest{}, 0, err
	}
	return v, len(data), err
}
