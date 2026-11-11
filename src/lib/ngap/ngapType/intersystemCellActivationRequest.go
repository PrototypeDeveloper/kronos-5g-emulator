// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntersystemCellActivationRequest represents ASN.1 IntersystemCellActivationRequest.
type IntersystemCellActivationRequest struct {
	ActivationID        int64 `aper:"valueExt,valueLB:0,valueUB:16384"`
	CellsToActivateList CellsToActivateList
	HasIEExtensions     bool
	IEExtensions        []IntersystemCellActivationRequestExtIEsEntry
	UnknownExtensions   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemCellActivationRequest) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemCellActivationRequest", v)
}

// DecodeIntersystemCellActivationRequest decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemCellActivationRequest(data []byte) (IntersystemCellActivationRequest, int, error) {
	var v IntersystemCellActivationRequest
	err := ngapcodec.DecodeParameter("IntersystemCellActivationRequest", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemCellActivationRequest{}, 0, err
	}
	return v, len(data), err
}
