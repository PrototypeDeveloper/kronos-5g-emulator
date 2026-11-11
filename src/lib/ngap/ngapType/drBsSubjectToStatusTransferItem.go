// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DRBsSubjectToStatusTransferItem represents ASN.1 DRBsSubjectToStatusTransferItem.
type DRBsSubjectToStatusTransferItem struct {
	DRBID             DRBID
	DRBStatusUL       DRBStatusUL
	DRBStatusDL       DRBStatusDL
	HasIEExtension    bool
	IEExtension       []DRBsSubjectToStatusTransferItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBsSubjectToStatusTransferItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DRBsSubjectToStatusTransferItem", v)
}

// DecodeDRBsSubjectToStatusTransferItem decodes a standalone APER value and reports the consumed octets.
func DecodeDRBsSubjectToStatusTransferItem(data []byte) (DRBsSubjectToStatusTransferItem, int, error) {
	var v DRBsSubjectToStatusTransferItem
	err := ngapcodec.DecodeParameter("DRBsSubjectToStatusTransferItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBsSubjectToStatusTransferItem{}, 0, err
	}
	return v, len(data), err
}
