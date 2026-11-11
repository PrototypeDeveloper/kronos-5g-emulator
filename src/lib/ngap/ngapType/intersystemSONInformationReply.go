// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntersystemSONInformationReply represents ASN.1 IntersystemSONInformationReply.
type IntersystemSONInformationReplyPresent uint32

const (
	IntersystemSONInformationReplyPresentNothing IntersystemSONInformationReplyPresent = iota
	IntersystemSONInformationReplyPresentNGRANCellActivation
	IntersystemSONInformationReplyPresentResourceStatus
	IntersystemSONInformationReplyPresentChoiceExtensions
)

type IntersystemSONInformationReply struct {
	Present             IntersystemSONInformationReplyPresent
	NGRANCellActivation IntersystemCellActivationReply
	ResourceStatus      IntersystemResourceStatusReply
	ChoiceExtensions    IntersystemSONInformationReplyExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemSONInformationReply) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemSONInformationReply", v)
}

// DecodeIntersystemSONInformationReply decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemSONInformationReply(data []byte) (IntersystemSONInformationReply, int, error) {
	var v IntersystemSONInformationReply
	err := ngapcodec.DecodeParameter("IntersystemSONInformationReply", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemSONInformationReply{}, 0, err
	}
	return v, len(data), err
}
