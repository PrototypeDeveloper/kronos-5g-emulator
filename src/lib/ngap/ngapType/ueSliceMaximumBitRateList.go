// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UESliceMaximumBitRateList represents ASN.1 UESliceMaximumBitRateList.
type UESliceMaximumBitRateList struct {
	List []UESliceMaximumBitRateItem `aper:"sizeLB:1,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UESliceMaximumBitRateList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UESliceMaximumBitRateList", v)
}

// DecodeUESliceMaximumBitRateList decodes a standalone APER value and reports the consumed octets.
func DecodeUESliceMaximumBitRateList(data []byte) (UESliceMaximumBitRateList, int, error) {
	var v UESliceMaximumBitRateList
	err := ngapcodec.DecodeParameter("UESliceMaximumBitRateList", data, &v)
	if ngapcodec.IsFatal(err) {
		return UESliceMaximumBitRateList{}, 0, err
	}
	return v, len(data), err
}
