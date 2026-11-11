// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntersystemSONInformationRequest represents ASN.1 IntersystemSONInformationRequest.
type IntersystemSONInformationRequestPresent uint32

const (
	IntersystemSONInformationRequestPresentNothing IntersystemSONInformationRequestPresent = iota
	IntersystemSONInformationRequestPresentNGRANCellActivation
	IntersystemSONInformationRequestPresentResourceStatus
	IntersystemSONInformationRequestPresentChoiceExtensions
)

type IntersystemSONInformationRequest struct {
	Present             IntersystemSONInformationRequestPresent
	NGRANCellActivation IntersystemCellActivationRequest
	ResourceStatus      IntersystemResourceStatusRequest
	ChoiceExtensions    IntersystemSONInformationRequestExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemSONInformationRequest) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemSONInformationRequest", v)
}

// DecodeIntersystemSONInformationRequest decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemSONInformationRequest(data []byte) (IntersystemSONInformationRequest, int, error) {
	var v IntersystemSONInformationRequest
	err := ngapcodec.DecodeParameter("IntersystemSONInformationRequest", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemSONInformationRequest{}, 0, err
	}
	return v, len(data), err
}
