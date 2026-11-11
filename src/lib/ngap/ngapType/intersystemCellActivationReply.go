// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntersystemCellActivationReply represents ASN.1 IntersystemCellActivationReply.
type IntersystemCellActivationReply struct {
	ActivatedCellList ActivatedCellList
	ActivationID      int64 `aper:"valueExt,valueLB:0,valueUB:16384"`
	HasIEExtensions   bool
	IEExtensions      []IntersystemCellActivationReplyExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemCellActivationReply) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemCellActivationReply", v)
}

// DecodeIntersystemCellActivationReply decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemCellActivationReply(data []byte) (IntersystemCellActivationReply, int, error) {
	var v IntersystemCellActivationReply
	err := ngapcodec.DecodeParameter("IntersystemCellActivationReply", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemCellActivationReply{}, 0, err
	}
	return v, len(data), err
}
