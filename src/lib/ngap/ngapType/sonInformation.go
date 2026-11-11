// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SONInformation represents ASN.1 SONInformation.
type SONInformationPresent uint32

const (
	SONInformationPresentNothing SONInformationPresent = iota
	SONInformationPresentSONInformationRequest
	SONInformationPresentSONInformationReply
	SONInformationPresentChoiceExtensions
)

type SONInformation struct {
	Present               SONInformationPresent
	SONInformationRequest SONInformationRequest
	SONInformationReply   SONInformationReply
	ChoiceExtensions      SONInformationExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SONInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SONInformation", v)
}

// DecodeSONInformation decodes a standalone APER value and reports the consumed octets.
func DecodeSONInformation(data []byte) (SONInformation, int, error) {
	var v SONInformation
	err := ngapcodec.DecodeParameter("SONInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return SONInformation{}, 0, err
	}
	return v, len(data), err
}
