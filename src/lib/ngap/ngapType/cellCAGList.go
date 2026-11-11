// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellCAGList represents ASN.1 CellCAGList.
type CellCAGList struct {
	List []CAGID `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellCAGList) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("CellCAGList", v) }

// DecodeCellCAGList decodes a standalone APER value and reports the consumed octets.
func DecodeCellCAGList(data []byte) (CellCAGList, int, error) {
	var v CellCAGList
	err := ngapcodec.DecodeParameter("CellCAGList", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellCAGList{}, 0, err
	}
	return v, len(data), err
}
