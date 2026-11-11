// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellsToActivateList represents ASN.1 CellsToActivateList.
type CellsToActivateList struct {
	List []NGRANCGI `aper:"sizeLB:1,sizeUB:16384"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellsToActivateList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellsToActivateList", v)
}

// DecodeCellsToActivateList decodes a standalone APER value and reports the consumed octets.
func DecodeCellsToActivateList(data []byte) (CellsToActivateList, int, error) {
	var v CellsToActivateList
	err := ngapcodec.DecodeParameter("CellsToActivateList", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellsToActivateList{}, 0, err
	}
	return v, len(data), err
}
