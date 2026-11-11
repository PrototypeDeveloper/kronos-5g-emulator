// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIoTSecurityInfo represents ASN.1 AIoT-SecurityInfo.
type AIoTSecurityInfoPresent uint32

const (
	AIoTSecurityInfoPresentNothing AIoTSecurityInfoPresent = iota
	AIoTSecurityInfoPresentAIoTRANDN
	AIoTSecurityInfoPresentChoiceExtensions
)

type AIoTSecurityInfo struct {
	Present          AIoTSecurityInfoPresent
	AIoTRANDN        BitString `aper:"sizeExt,sizeLB:128,sizeUB:128"`
	ChoiceExtensions AIoTSecurityInfoExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTSecurityInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-SecurityInfo", v)
}

// DecodeAIoTSecurityInfo decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTSecurityInfo(data []byte) (AIoTSecurityInfo, int, error) {
	var v AIoTSecurityInfo
	err := ngapcodec.DecodeParameter("AIoT-SecurityInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTSecurityInfo{}, 0, err
	}
	return v, len(data), err
}
