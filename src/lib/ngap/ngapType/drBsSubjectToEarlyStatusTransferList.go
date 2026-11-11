// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DRBsSubjectToEarlyStatusTransferList represents ASN.1 DRBsSubjectToEarlyStatusTransfer-List.
type DRBsSubjectToEarlyStatusTransferList struct {
	List []DRBsSubjectToEarlyStatusTransferItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBsSubjectToEarlyStatusTransferList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DRBsSubjectToEarlyStatusTransfer-List", v)
}

// DecodeDRBsSubjectToEarlyStatusTransferList decodes a standalone APER value and reports the consumed octets.
func DecodeDRBsSubjectToEarlyStatusTransferList(data []byte) (DRBsSubjectToEarlyStatusTransferList, int, error) {
	var v DRBsSubjectToEarlyStatusTransferList
	err := ngapcodec.DecodeParameter("DRBsSubjectToEarlyStatusTransfer-List", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBsSubjectToEarlyStatusTransferList{}, 0, err
	}
	return v, len(data), err
}
