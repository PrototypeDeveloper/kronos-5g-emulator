// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NPNPagingAssistanceInformation represents ASN.1 NPN-PagingAssistanceInformation.
type NPNPagingAssistanceInformationPresent uint32

const (
	NPNPagingAssistanceInformationPresentNothing NPNPagingAssistanceInformationPresent = iota
	NPNPagingAssistanceInformationPresentPNINPNPagingAssistance
	NPNPagingAssistanceInformationPresentChoiceExtensions
)

type NPNPagingAssistanceInformation struct {
	Present                NPNPagingAssistanceInformationPresent
	PNINPNPagingAssistance AllowedPNINPNList
	ChoiceExtensions       NPNPagingAssistanceInformationExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NPNPagingAssistanceInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NPN-PagingAssistanceInformation", v)
}

// DecodeNPNPagingAssistanceInformation decodes a standalone APER value and reports the consumed octets.
func DecodeNPNPagingAssistanceInformation(data []byte) (NPNPagingAssistanceInformation, int, error) {
	var v NPNPagingAssistanceInformation
	err := ngapcodec.DecodeParameter("NPN-PagingAssistanceInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return NPNPagingAssistanceInformation{}, 0, err
	}
	return v, len(data), err
}
