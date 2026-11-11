// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRNTNTAIInformation represents ASN.1 NRNTNTAIInformation.
type NRNTNTAIInformation struct {
	ServingPLMN                    PLMNIdentity
	TACListInNRNTN                 TACListInNRNTN
	HasUELocationDerivedTACInNRNTN bool
	UELocationDerivedTACInNRNTN    TAC
	HasIEExtensions                bool
	IEExtensions                   []NRNTNTAIInformationExtIEsEntry
	UnknownExtensions              []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRNTNTAIInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NRNTNTAIInformation", v)
}

// DecodeNRNTNTAIInformation decodes a standalone APER value and reports the consumed octets.
func DecodeNRNTNTAIInformation(data []byte) (NRNTNTAIInformation, int, error) {
	var v NRNTNTAIInformation
	err := ngapcodec.DecodeParameter("NRNTNTAIInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRNTNTAIInformation{}, 0, err
	}
	return v, len(data), err
}
