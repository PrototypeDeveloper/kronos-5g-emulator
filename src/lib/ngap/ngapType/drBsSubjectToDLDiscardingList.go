// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DRBsSubjectToDLDiscardingList represents ASN.1 DRBsSubjectToDLDiscarding-List.
type DRBsSubjectToDLDiscardingList struct {
	List []DRBsSubjectToDLDiscardingItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBsSubjectToDLDiscardingList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DRBsSubjectToDLDiscarding-List", v)
}

// DecodeDRBsSubjectToDLDiscardingList decodes a standalone APER value and reports the consumed octets.
func DecodeDRBsSubjectToDLDiscardingList(data []byte) (DRBsSubjectToDLDiscardingList, int, error) {
	var v DRBsSubjectToDLDiscardingList
	err := ngapcodec.DecodeParameter("DRBsSubjectToDLDiscarding-List", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBsSubjectToDLDiscardingList{}, 0, err
	}
	return v, len(data), err
}
