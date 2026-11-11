// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionForPagingItem represents ASN.1 PDUSessionForPagingItem.
type PDUSessionForPagingItem struct {
	PDUSessionID                    PDUSessionID
	PagingPolicyDifferentiationList PagingPolicyDifferentiationList
	HasIEExtensions                 bool
	IEExtensions                    []PDUSessionForPagingItemExtIEsEntry
	UnknownExtensions               []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionForPagingItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionForPagingItem", v)
}

// DecodePDUSessionForPagingItem decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionForPagingItem(data []byte) (PDUSessionForPagingItem, int, error) {
	var v PDUSessionForPagingItem
	err := ngapcodec.DecodeParameter("PDUSessionForPagingItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionForPagingItem{}, 0, err
	}
	return v, len(data), err
}
