// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntersystemResourceStatusRequest represents ASN.1 IntersystemResourceStatusRequest.
type IntersystemResourceStatusRequest struct {
	ReportingSystem       ReportingSystem
	ReportCharacteristics ReportCharacteristics
	ReportType            ReportType
	HasIEExtensions       bool
	IEExtensions          []IntersystemResourceStatusRequestExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemResourceStatusRequest) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemResourceStatusRequest", v)
}

// DecodeIntersystemResourceStatusRequest decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemResourceStatusRequest(data []byte) (IntersystemResourceStatusRequest, int, error) {
	var v IntersystemResourceStatusRequest
	err := ngapcodec.DecodeParameter("IntersystemResourceStatusRequest", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemResourceStatusRequest{}, 0, err
	}
	return v, len(data), err
}
