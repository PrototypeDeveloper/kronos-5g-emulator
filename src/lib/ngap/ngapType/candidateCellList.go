// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CandidateCellList represents ASN.1 CandidateCellList.
type CandidateCellList struct {
	List []CandidateCellItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CandidateCellList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CandidateCellList", v)
}

// DecodeCandidateCellList decodes a standalone APER value and reports the consumed octets.
func DecodeCandidateCellList(data []byte) (CandidateCellList, int, error) {
	var v CandidateCellList
	err := ngapcodec.DecodeParameter("CandidateCellList", data, &v)
	if ngapcodec.IsFatal(err) {
		return CandidateCellList{}, 0, err
	}
	return v, len(data), err
}
