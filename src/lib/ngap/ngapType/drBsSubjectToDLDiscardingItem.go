// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DRBsSubjectToDLDiscardingItem represents ASN.1 DRBsSubjectToDLDiscarding-Item.
type DRBsSubjectToDLDiscardingItem struct {
	DrbID             DRBID
	DlCount           DLCountChoice
	HasIEExtension    bool
	IEExtension       []DRBsSubjectToDLDiscardingItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBsSubjectToDLDiscardingItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DRBsSubjectToDLDiscarding-Item", v)
}

// DecodeDRBsSubjectToDLDiscardingItem decodes a standalone APER value and reports the consumed octets.
func DecodeDRBsSubjectToDLDiscardingItem(data []byte) (DRBsSubjectToDLDiscardingItem, int, error) {
	var v DRBsSubjectToDLDiscardingItem
	err := ngapcodec.DecodeParameter("DRBsSubjectToDLDiscarding-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBsSubjectToDLDiscardingItem{}, 0, err
	}
	return v, len(data), err
}
