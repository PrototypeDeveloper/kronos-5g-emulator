// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LPWUSPSAssistanceInformation represents ASN.1 LPWUSPSAssistanceInformation.
type LPWUSPSAssistanceInformation struct {
	LPWUScNsubgroupID LPWUSCNsubgroupID
	HasIEExtensions   bool
	IEExtensions      []LPWUSPSAssistanceInformationExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LPWUSPSAssistanceInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LPWUSPSAssistanceInformation", v)
}

// DecodeLPWUSPSAssistanceInformation decodes a standalone APER value and reports the consumed octets.
func DecodeLPWUSPSAssistanceInformation(data []byte) (LPWUSPSAssistanceInformation, int, error) {
	var v LPWUSPSAssistanceInformation
	err := ngapcodec.DecodeParameter("LPWUSPSAssistanceInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return LPWUSPSAssistanceInformation{}, 0, err
	}
	return v, len(data), err
}
