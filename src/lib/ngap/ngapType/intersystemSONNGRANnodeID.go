// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntersystemSONNGRANnodeID represents ASN.1 IntersystemSONNGRANnodeID.
type IntersystemSONNGRANnodeID struct {
	GlobalRANNodeID   GlobalRANNodeID
	SelectedTAI       TAI
	HasIEExtensions   bool
	IEExtensions      []IntersystemSONNGRANnodeIDExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemSONNGRANnodeID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemSONNGRANnodeID", v)
}

// DecodeIntersystemSONNGRANnodeID decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemSONNGRANnodeID(data []byte) (IntersystemSONNGRANnodeID, int, error) {
	var v IntersystemSONNGRANnodeID
	err := ngapcodec.DecodeParameter("IntersystemSONNGRANnodeID", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemSONNGRANnodeID{}, 0, err
	}
	return v, len(data), err
}
