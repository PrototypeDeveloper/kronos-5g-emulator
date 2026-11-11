// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// WUSAssistanceInformation represents ASN.1 WUS-Assistance-Information.
type WUSAssistanceInformation struct {
	PagingProbabilityInformation PagingProbabilityInformation
	HasIEExtensions              bool
	IEExtensions                 []WUSAssistanceInformationExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v WUSAssistanceInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("WUS-Assistance-Information", v)
}

// DecodeWUSAssistanceInformation decodes a standalone APER value and reports the consumed octets.
func DecodeWUSAssistanceInformation(data []byte) (WUSAssistanceInformation, int, error) {
	var v WUSAssistanceInformation
	err := ngapcodec.DecodeParameter("WUS-Assistance-Information", data, &v)
	if ngapcodec.IsFatal(err) {
		return WUSAssistanceInformation{}, 0, err
	}
	return v, len(data), err
}
