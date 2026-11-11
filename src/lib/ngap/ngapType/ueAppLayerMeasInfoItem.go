// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEAppLayerMeasInfoItem represents ASN.1 UEAppLayerMeasInfoItem.
type UEAppLayerMeasInfoItem struct {
	UEAppLayerMeasConfigInfo UEAppLayerMeasConfigInfo
	HasIEExtensions          bool
	IEExtensions             []UEAppLayerMeasInfoItemExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEAppLayerMeasInfoItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEAppLayerMeasInfoItem", v)
}

// DecodeUEAppLayerMeasInfoItem decodes a standalone APER value and reports the consumed octets.
func DecodeUEAppLayerMeasInfoItem(data []byte) (UEAppLayerMeasInfoItem, int, error) {
	var v UEAppLayerMeasInfoItem
	err := ngapcodec.DecodeParameter("UEAppLayerMeasInfoItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEAppLayerMeasInfoItem{}, 0, err
	}
	return v, len(data), err
}
