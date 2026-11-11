// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AMFTNLAssociationToRemoveList represents ASN.1 AMF-TNLAssociationToRemoveList.
type AMFTNLAssociationToRemoveList struct {
	List []AMFTNLAssociationToRemoveItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFTNLAssociationToRemoveList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AMF-TNLAssociationToRemoveList", v)
}

// DecodeAMFTNLAssociationToRemoveList decodes a standalone APER value and reports the consumed octets.
func DecodeAMFTNLAssociationToRemoveList(data []byte) (AMFTNLAssociationToRemoveList, int, error) {
	var v AMFTNLAssociationToRemoveList
	err := ngapcodec.DecodeParameter("AMF-TNLAssociationToRemoveList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFTNLAssociationToRemoveList{}, 0, err
	}
	return v, len(data), err
}
