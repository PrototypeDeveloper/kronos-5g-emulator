// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SliceSupportList represents ASN.1 SliceSupportList.
type SliceSupportList struct {
	List []SliceSupportItem `aper:"sizeLB:1,sizeUB:1024"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SliceSupportList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SliceSupportList", v)
}

// DecodeSliceSupportList decodes a standalone APER value and reports the consumed octets.
func DecodeSliceSupportList(data []byte) (SliceSupportList, int, error) {
	var v SliceSupportList
	err := ngapcodec.DecodeParameter("SliceSupportList", data, &v)
	if ngapcodec.IsFatal(err) {
		return SliceSupportList{}, 0, err
	}
	return v, len(data), err
}
