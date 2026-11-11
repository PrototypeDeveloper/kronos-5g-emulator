// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntersystemResourceStatusReply represents ASN.1 IntersystemResourceStatusReply.
type IntersystemResourceStatusReply struct {
	Reportingsystem   ReportingSystem
	HasIEExtensions   bool
	IEExtensions      []IntersystemResourceStatusReplyExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemResourceStatusReply) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemResourceStatusReply", v)
}

// DecodeIntersystemResourceStatusReply decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemResourceStatusReply(data []byte) (IntersystemResourceStatusReply, int, error) {
	var v IntersystemResourceStatusReply
	err := ngapcodec.DecodeParameter("IntersystemResourceStatusReply", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemResourceStatusReply{}, 0, err
	}
	return v, len(data), err
}
