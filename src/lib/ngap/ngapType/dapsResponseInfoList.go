// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DAPSResponseInfoList represents ASN.1 DAPSResponseInfoList.
type DAPSResponseInfoList struct {
	List []DAPSResponseInfoItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DAPSResponseInfoList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DAPSResponseInfoList", v)
}

// DecodeDAPSResponseInfoList decodes a standalone APER value and reports the consumed octets.
func DecodeDAPSResponseInfoList(data []byte) (DAPSResponseInfoList, int, error) {
	var v DAPSResponseInfoList
	err := ngapcodec.DecodeParameter("DAPSResponseInfoList", data, &v)
	if ngapcodec.IsFatal(err) {
		return DAPSResponseInfoList{}, 0, err
	}
	return v, len(data), err
}
