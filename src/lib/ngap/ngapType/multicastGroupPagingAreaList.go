// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MulticastGroupPagingAreaList represents ASN.1 MulticastGroupPagingAreaList.
type MulticastGroupPagingAreaList struct {
	List []MulticastGroupPagingAreaItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MulticastGroupPagingAreaList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MulticastGroupPagingAreaList", v)
}

// DecodeMulticastGroupPagingAreaList decodes a standalone APER value and reports the consumed octets.
func DecodeMulticastGroupPagingAreaList(data []byte) (MulticastGroupPagingAreaList, int, error) {
	var v MulticastGroupPagingAreaList
	err := ngapcodec.DecodeParameter("MulticastGroupPagingAreaList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MulticastGroupPagingAreaList{}, 0, err
	}
	return v, len(data), err
}
