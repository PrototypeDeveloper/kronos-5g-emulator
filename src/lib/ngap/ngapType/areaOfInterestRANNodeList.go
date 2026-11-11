// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AreaOfInterestRANNodeList represents ASN.1 AreaOfInterestRANNodeList.
type AreaOfInterestRANNodeList struct {
	List []AreaOfInterestRANNodeItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaOfInterestRANNodeList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaOfInterestRANNodeList", v)
}

// DecodeAreaOfInterestRANNodeList decodes a standalone APER value and reports the consumed octets.
func DecodeAreaOfInterestRANNodeList(data []byte) (AreaOfInterestRANNodeList, int, error) {
	var v AreaOfInterestRANNodeList
	err := ngapcodec.DecodeParameter("AreaOfInterestRANNodeList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaOfInterestRANNodeList{}, 0, err
	}
	return v, len(data), err
}
