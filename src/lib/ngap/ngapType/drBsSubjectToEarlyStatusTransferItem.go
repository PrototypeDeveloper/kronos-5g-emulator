// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DRBsSubjectToEarlyStatusTransferItem represents ASN.1 DRBsSubjectToEarlyStatusTransfer-Item.
type DRBsSubjectToEarlyStatusTransferItem struct {
	DRBID             DRBID
	FirstDLCOUNT      DRBStatusDL
	HasIEExtension    bool
	IEExtension       []DRBsSubjectToEarlyStatusTransferItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBsSubjectToEarlyStatusTransferItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DRBsSubjectToEarlyStatusTransfer-Item", v)
}

// DecodeDRBsSubjectToEarlyStatusTransferItem decodes a standalone APER value and reports the consumed octets.
func DecodeDRBsSubjectToEarlyStatusTransferItem(data []byte) (DRBsSubjectToEarlyStatusTransferItem, int, error) {
	var v DRBsSubjectToEarlyStatusTransferItem
	err := ngapcodec.DecodeParameter("DRBsSubjectToEarlyStatusTransfer-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBsSubjectToEarlyStatusTransferItem{}, 0, err
	}
	return v, len(data), err
}
