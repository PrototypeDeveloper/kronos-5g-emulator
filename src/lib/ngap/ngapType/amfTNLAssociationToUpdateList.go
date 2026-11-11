// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AMFTNLAssociationToUpdateList represents ASN.1 AMF-TNLAssociationToUpdateList.
type AMFTNLAssociationToUpdateList struct {
	List []AMFTNLAssociationToUpdateItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFTNLAssociationToUpdateList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AMF-TNLAssociationToUpdateList", v)
}

// DecodeAMFTNLAssociationToUpdateList decodes a standalone APER value and reports the consumed octets.
func DecodeAMFTNLAssociationToUpdateList(data []byte) (AMFTNLAssociationToUpdateList, int, error) {
	var v AMFTNLAssociationToUpdateList
	err := ngapcodec.DecodeParameter("AMF-TNLAssociationToUpdateList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFTNLAssociationToUpdateList{}, 0, err
	}
	return v, len(data), err
}
