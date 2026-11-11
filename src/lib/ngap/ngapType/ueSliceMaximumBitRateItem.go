// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UESliceMaximumBitRateItem represents ASN.1 UESliceMaximumBitRateItem.
type UESliceMaximumBitRateItem struct {
	SNSSAI                  SNSSAI
	UESliceMaximumBitRateDL BitRate
	UESliceMaximumBitRateUL BitRate
	HasIEExtensions         bool
	IEExtensions            []UESliceMaximumBitRateItemExtIEsEntry
	UnknownExtensions       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UESliceMaximumBitRateItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UESliceMaximumBitRateItem", v)
}

// DecodeUESliceMaximumBitRateItem decodes a standalone APER value and reports the consumed octets.
func DecodeUESliceMaximumBitRateItem(data []byte) (UESliceMaximumBitRateItem, int, error) {
	var v UESliceMaximumBitRateItem
	err := ngapcodec.DecodeParameter("UESliceMaximumBitRateItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return UESliceMaximumBitRateItem{}, 0, err
	}
	return v, len(data), err
}
