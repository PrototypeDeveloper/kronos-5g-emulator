// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ULNGUUPTNLModifyList represents ASN.1 UL-NGU-UP-TNLModifyList.
type ULNGUUPTNLModifyList struct {
	List []ULNGUUPTNLModifyItem `aper:"sizeLB:1,sizeUB:4"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ULNGUUPTNLModifyList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UL-NGU-UP-TNLModifyList", v)
}

// DecodeULNGUUPTNLModifyList decodes a standalone APER value and reports the consumed octets.
func DecodeULNGUUPTNLModifyList(data []byte) (ULNGUUPTNLModifyList, int, error) {
	var v ULNGUUPTNLModifyList
	err := ngapcodec.DecodeParameter("UL-NGU-UP-TNLModifyList", data, &v)
	if ngapcodec.IsFatal(err) {
		return ULNGUUPTNLModifyList{}, 0, err
	}
	return v, len(data), err
}
