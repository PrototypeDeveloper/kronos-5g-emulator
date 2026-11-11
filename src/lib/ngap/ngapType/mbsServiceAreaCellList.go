// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSServiceAreaCellList represents ASN.1 MBS-ServiceAreaCellList.
type MBSServiceAreaCellList struct {
	List []NRCGI `aper:"sizeLB:1,sizeUB:8192"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSServiceAreaCellList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-ServiceAreaCellList", v)
}

// DecodeMBSServiceAreaCellList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSServiceAreaCellList(data []byte) (MBSServiceAreaCellList, int, error) {
	var v MBSServiceAreaCellList
	err := ngapcodec.DecodeParameter("MBS-ServiceAreaCellList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSServiceAreaCellList{}, 0, err
	}
	return v, len(data), err
}
