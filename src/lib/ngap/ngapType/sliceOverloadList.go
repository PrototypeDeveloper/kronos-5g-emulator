// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SliceOverloadList represents ASN.1 SliceOverloadList.
type SliceOverloadList struct {
	List []SliceOverloadItem `aper:"sizeLB:1,sizeUB:1024"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SliceOverloadList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SliceOverloadList", v)
}

// DecodeSliceOverloadList decodes a standalone APER value and reports the consumed octets.
func DecodeSliceOverloadList(data []byte) (SliceOverloadList, int, error) {
	var v SliceOverloadList
	err := ngapcodec.DecodeParameter("SliceOverloadList", data, &v)
	if ngapcodec.IsFatal(err) {
		return SliceOverloadList{}, 0, err
	}
	return v, len(data), err
}
