// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIoTCommandAssistanceInformation represents ASN.1 AIoT-CommandAssistanceInformation.
type AIoTCommandAssistanceInformation struct {
	EstimateofExpectedD2RMsgSize int64 `aper:"valueExt,valueLB:1,valueUB:256"`
	HasIEExtensions              bool
	IEExtensions                 []AIoTCommandAssistanceInformationExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTCommandAssistanceInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-CommandAssistanceInformation", v)
}

// DecodeAIoTCommandAssistanceInformation decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTCommandAssistanceInformation(data []byte) (AIoTCommandAssistanceInformation, int, error) {
	var v AIoTCommandAssistanceInformation
	err := ngapcodec.DecodeParameter("AIoT-CommandAssistanceInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTCommandAssistanceInformation{}, 0, err
	}
	return v, len(data), err
}
