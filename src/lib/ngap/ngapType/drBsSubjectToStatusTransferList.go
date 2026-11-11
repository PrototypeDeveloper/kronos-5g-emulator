// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DRBsSubjectToStatusTransferList represents ASN.1 DRBsSubjectToStatusTransferList.
type DRBsSubjectToStatusTransferList struct {
	List []DRBsSubjectToStatusTransferItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBsSubjectToStatusTransferList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DRBsSubjectToStatusTransferList", v)
}

// DecodeDRBsSubjectToStatusTransferList decodes a standalone APER value and reports the consumed octets.
func DecodeDRBsSubjectToStatusTransferList(data []byte) (DRBsSubjectToStatusTransferList, int, error) {
	var v DRBsSubjectToStatusTransferList
	err := ngapcodec.DecodeParameter("DRBsSubjectToStatusTransferList", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBsSubjectToStatusTransferList{}, 0, err
	}
	return v, len(data), err
}
