// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MulticastGroupPagingAreaItem represents ASN.1 MulticastGroupPagingAreaItem.
type MulticastGroupPagingAreaItem struct {
	MulticastGroupPagingArea MulticastGroupPagingArea
	HasUEPagingList          bool
	UEPagingList             UEPagingList
	HasIEExtensions          bool
	IEExtensions             []MulticastGroupPagingAreaItemExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MulticastGroupPagingAreaItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MulticastGroupPagingAreaItem", v)
}

// DecodeMulticastGroupPagingAreaItem decodes a standalone APER value and reports the consumed octets.
func DecodeMulticastGroupPagingAreaItem(data []byte) (MulticastGroupPagingAreaItem, int, error) {
	var v MulticastGroupPagingAreaItem
	err := ngapcodec.DecodeParameter("MulticastGroupPagingAreaItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return MulticastGroupPagingAreaItem{}, 0, err
	}
	return v, len(data), err
}
