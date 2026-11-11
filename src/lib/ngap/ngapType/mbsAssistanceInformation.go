// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSAssistanceInformation represents ASN.1 MBS-AssistanceInformation.
const (
	MBSAssistanceInformationPresentTrue Enumerated = 0
)

type MBSAssistanceInformation struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSAssistanceInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-AssistanceInformation", v)
}

// DecodeMBSAssistanceInformation decodes a standalone APER value and reports the consumed octets.
func DecodeMBSAssistanceInformation(data []byte) (MBSAssistanceInformation, int, error) {
	var v MBSAssistanceInformation
	err := ngapcodec.DecodeParameter("MBS-AssistanceInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSAssistanceInformation{}, 0, err
	}
	return v, len(data), err
}
