// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEPagingList represents ASN.1 UE-PagingList.
type UEPagingList struct {
	List []UEPagingItem `aper:"sizeLB:1,sizeUB:4096"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEPagingList) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("UE-PagingList", v) }

// DecodeUEPagingList decodes a standalone APER value and reports the consumed octets.
func DecodeUEPagingList(data []byte) (UEPagingList, int, error) {
	var v UEPagingList
	err := ngapcodec.DecodeParameter("UE-PagingList", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEPagingList{}, 0, err
	}
	return v, len(data), err
}
