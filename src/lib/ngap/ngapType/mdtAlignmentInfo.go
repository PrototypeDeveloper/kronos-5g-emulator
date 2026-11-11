// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MDTAlignmentInfo represents ASN.1 MDT-AlignmentInfo.
type MDTAlignmentInfoPresent uint32

const (
	MDTAlignmentInfoPresentNothing MDTAlignmentInfoPresent = iota
	MDTAlignmentInfoPresentSBasedMDT
	MDTAlignmentInfoPresentChoiceExtensions
)

type MDTAlignmentInfo struct {
	Present          MDTAlignmentInfoPresent
	SBasedMDT        NGRANTraceID
	ChoiceExtensions MDTAlignmentInfoExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTAlignmentInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MDT-AlignmentInfo", v)
}

// DecodeMDTAlignmentInfo decodes a standalone APER value and reports the consumed octets.
func DecodeMDTAlignmentInfo(data []byte) (MDTAlignmentInfo, int, error) {
	var v MDTAlignmentInfo
	err := ngapcodec.DecodeParameter("MDT-AlignmentInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTAlignmentInfo{}, 0, err
	}
	return v, len(data), err
}
