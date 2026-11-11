// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// OverloadResponse represents ASN.1 OverloadResponse.
type OverloadResponsePresent uint32

const (
	OverloadResponsePresentNothing OverloadResponsePresent = iota
	OverloadResponsePresentOverloadAction
	OverloadResponsePresentChoiceExtensions
)

type OverloadResponse struct {
	Present          OverloadResponsePresent
	OverloadAction   OverloadAction
	ChoiceExtensions OverloadResponseExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v OverloadResponse) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("OverloadResponse", v)
}

// DecodeOverloadResponse decodes a standalone APER value and reports the consumed octets.
func DecodeOverloadResponse(data []byte) (OverloadResponse, int, error) {
	var v OverloadResponse
	err := ngapcodec.DecodeParameter("OverloadResponse", data, &v)
	if ngapcodec.IsFatal(err) {
		return OverloadResponse{}, 0, err
	}
	return v, len(data), err
}
