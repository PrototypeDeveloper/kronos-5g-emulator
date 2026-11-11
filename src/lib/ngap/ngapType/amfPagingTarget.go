// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AMFPagingTarget represents ASN.1 AMFPagingTarget.
type AMFPagingTargetPresent uint32

const (
	AMFPagingTargetPresentNothing AMFPagingTargetPresent = iota
	AMFPagingTargetPresentGlobalRANNodeID
	AMFPagingTargetPresentTAI
	AMFPagingTargetPresentChoiceExtensions
)

type AMFPagingTarget struct {
	Present          AMFPagingTargetPresent
	GlobalRANNodeID  GlobalRANNodeID
	TAI              TAI
	ChoiceExtensions AMFPagingTargetExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFPagingTarget) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AMFPagingTarget", v)
}

// DecodeAMFPagingTarget decodes a standalone APER value and reports the consumed octets.
func DecodeAMFPagingTarget(data []byte) (AMFPagingTarget, int, error) {
	var v AMFPagingTarget
	err := ngapcodec.DecodeParameter("AMFPagingTarget", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFPagingTarget{}, 0, err
	}
	return v, len(data), err
}
