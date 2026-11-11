// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AMFTNLAssociationSetupList represents ASN.1 AMF-TNLAssociationSetupList.
type AMFTNLAssociationSetupList struct {
	List []AMFTNLAssociationSetupItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFTNLAssociationSetupList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AMF-TNLAssociationSetupList", v)
}

// DecodeAMFTNLAssociationSetupList decodes a standalone APER value and reports the consumed octets.
func DecodeAMFTNLAssociationSetupList(data []byte) (AMFTNLAssociationSetupList, int, error) {
	var v AMFTNLAssociationSetupList
	err := ngapcodec.DecodeParameter("AMF-TNLAssociationSetupList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFTNLAssociationSetupList{}, 0, err
	}
	return v, len(data), err
}
