// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ActivatedCellList represents ASN.1 ActivatedCellList.
type ActivatedCellList struct {
	List []NGRANCGI `aper:"sizeLB:1,sizeUB:16384"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ActivatedCellList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ActivatedCellList", v)
}

// DecodeActivatedCellList decodes a standalone APER value and reports the consumed octets.
func DecodeActivatedCellList(data []byte) (ActivatedCellList, int, error) {
	var v ActivatedCellList
	err := ngapcodec.DecodeParameter("ActivatedCellList", data, &v)
	if ngapcodec.IsFatal(err) {
		return ActivatedCellList{}, 0, err
	}
	return v, len(data), err
}
