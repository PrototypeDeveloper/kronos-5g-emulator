// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ULNGUUPTNLModifyItem represents ASN.1 UL-NGU-UP-TNLModifyItem.
type ULNGUUPTNLModifyItem struct {
	ULNGUUPTNLInformation UPTransportLayerInformation
	DLNGUUPTNLInformation UPTransportLayerInformation
	HasIEExtensions       bool
	IEExtensions          []ULNGUUPTNLModifyItemExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ULNGUUPTNLModifyItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UL-NGU-UP-TNLModifyItem", v)
}

// DecodeULNGUUPTNLModifyItem decodes a standalone APER value and reports the consumed octets.
func DecodeULNGUUPTNLModifyItem(data []byte) (ULNGUUPTNLModifyItem, int, error) {
	var v ULNGUUPTNLModifyItem
	err := ngapcodec.DecodeParameter("UL-NGU-UP-TNLModifyItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return ULNGUUPTNLModifyItem{}, 0, err
	}
	return v, len(data), err
}
