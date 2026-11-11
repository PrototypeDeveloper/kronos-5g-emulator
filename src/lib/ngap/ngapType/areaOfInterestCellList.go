// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AreaOfInterestCellList represents ASN.1 AreaOfInterestCellList.
type AreaOfInterestCellList struct {
	List []AreaOfInterestCellItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaOfInterestCellList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaOfInterestCellList", v)
}

// DecodeAreaOfInterestCellList decodes a standalone APER value and reports the consumed octets.
func DecodeAreaOfInterestCellList(data []byte) (AreaOfInterestCellList, int, error) {
	var v AreaOfInterestCellList
	err := ngapcodec.DecodeParameter("AreaOfInterestCellList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaOfInterestCellList{}, 0, err
	}
	return v, len(data), err
}
