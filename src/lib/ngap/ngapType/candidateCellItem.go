// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CandidateCellItem represents ASN.1 CandidateCellItem.
type CandidateCellItem struct {
	CandidateCell     CandidateCell
	HasIEExtensions   bool
	IEExtensions      []CandidateCellItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CandidateCellItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CandidateCellItem", v)
}

// DecodeCandidateCellItem decodes a standalone APER value and reports the consumed octets.
func DecodeCandidateCellItem(data []byte) (CandidateCellItem, int, error) {
	var v CandidateCellItem
	err := ngapcodec.DecodeParameter("CandidateCellItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return CandidateCellItem{}, 0, err
	}
	return v, len(data), err
}
