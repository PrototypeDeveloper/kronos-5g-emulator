// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AMFTNLAssociationToAddList represents ASN.1 AMF-TNLAssociationToAddList.
type AMFTNLAssociationToAddList struct {
	List []AMFTNLAssociationToAddItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFTNLAssociationToAddList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AMF-TNLAssociationToAddList", v)
}

// DecodeAMFTNLAssociationToAddList decodes a standalone APER value and reports the consumed octets.
func DecodeAMFTNLAssociationToAddList(data []byte) (AMFTNLAssociationToAddList, int, error) {
	var v AMFTNLAssociationToAddList
	err := ngapcodec.DecodeParameter("AMF-TNLAssociationToAddList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFTNLAssociationToAddList{}, 0, err
	}
	return v, len(data), err
}
