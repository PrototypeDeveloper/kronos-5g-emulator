// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEPagingItem represents ASN.1 UE-PagingItem.
type UEPagingItem struct {
	UEIdentityIndexValue UEIdentityIndexValue
	HasPagingDRX         bool
	PagingDRX            PagingDRX
	HasIEExtensions      bool
	IEExtensions         []UEPagingItemExtIEsEntry
	UnknownExtensions    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEPagingItem) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("UE-PagingItem", v) }

// DecodeUEPagingItem decodes a standalone APER value and reports the consumed octets.
func DecodeUEPagingItem(data []byte) (UEPagingItem, int, error) {
	var v UEPagingItem
	err := ngapcodec.DecodeParameter("UE-PagingItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEPagingItem{}, 0, err
	}
	return v, len(data), err
}
