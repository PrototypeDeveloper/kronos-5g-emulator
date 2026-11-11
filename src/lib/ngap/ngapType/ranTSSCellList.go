// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANTSSCellList represents ASN.1 RANTSSCellList.
type RANTSSCellList struct {
	List []RANTSSCellItem `aper:"sizeLB:1,sizeUB:16384"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANTSSCellList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RANTSSCellList", v)
}

// DecodeRANTSSCellList decodes a standalone APER value and reports the consumed octets.
func DecodeRANTSSCellList(data []byte) (RANTSSCellList, int, error) {
	var v RANTSSCellList
	err := ngapcodec.DecodeParameter("RANTSSCellList", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANTSSCellList{}, 0, err
	}
	return v, len(data), err
}
