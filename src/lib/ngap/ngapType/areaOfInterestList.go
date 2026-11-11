// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AreaOfInterestList represents ASN.1 AreaOfInterestList.
type AreaOfInterestList struct {
	List []AreaOfInterestItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaOfInterestList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaOfInterestList", v)
}

// DecodeAreaOfInterestList decodes a standalone APER value and reports the consumed octets.
func DecodeAreaOfInterestList(data []byte) (AreaOfInterestList, int, error) {
	var v AreaOfInterestList
	err := ngapcodec.DecodeParameter("AreaOfInterestList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaOfInterestList{}, 0, err
	}
	return v, len(data), err
}
