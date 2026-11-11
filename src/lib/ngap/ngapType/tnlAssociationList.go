// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TNLAssociationList represents ASN.1 TNLAssociationList.
type TNLAssociationList struct {
	List []TNLAssociationItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TNLAssociationList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TNLAssociationList", v)
}

// DecodeTNLAssociationList decodes a standalone APER value and reports the consumed octets.
func DecodeTNLAssociationList(data []byte) (TNLAssociationList, int, error) {
	var v TNLAssociationList
	err := ngapcodec.DecodeParameter("TNLAssociationList", data, &v)
	if ngapcodec.IsFatal(err) {
		return TNLAssociationList{}, 0, err
	}
	return v, len(data), err
}
