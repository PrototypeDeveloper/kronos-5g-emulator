// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGRANTNLAssociationToRemoveList represents ASN.1 NGRAN-TNLAssociationToRemoveList.
type NGRANTNLAssociationToRemoveList struct {
	List []NGRANTNLAssociationToRemoveItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANTNLAssociationToRemoveList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGRAN-TNLAssociationToRemoveList", v)
}

// DecodeNGRANTNLAssociationToRemoveList decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANTNLAssociationToRemoveList(data []byte) (NGRANTNLAssociationToRemoveList, int, error) {
	var v NGRANTNLAssociationToRemoveList
	err := ngapcodec.DecodeParameter("NGRAN-TNLAssociationToRemoveList", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANTNLAssociationToRemoveList{}, 0, err
	}
	return v, len(data), err
}
