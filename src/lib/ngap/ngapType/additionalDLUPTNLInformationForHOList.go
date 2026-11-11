// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AdditionalDLUPTNLInformationForHOList represents ASN.1 AdditionalDLUPTNLInformationForHOList.
type AdditionalDLUPTNLInformationForHOList struct {
	List []AdditionalDLUPTNLInformationForHOItem `aper:"sizeLB:1,sizeUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AdditionalDLUPTNLInformationForHOList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AdditionalDLUPTNLInformationForHOList", v)
}

// DecodeAdditionalDLUPTNLInformationForHOList decodes a standalone APER value and reports the consumed octets.
func DecodeAdditionalDLUPTNLInformationForHOList(data []byte) (AdditionalDLUPTNLInformationForHOList, int, error) {
	var v AdditionalDLUPTNLInformationForHOList
	err := ngapcodec.DecodeParameter("AdditionalDLUPTNLInformationForHOList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AdditionalDLUPTNLInformationForHOList{}, 0, err
	}
	return v, len(data), err
}
