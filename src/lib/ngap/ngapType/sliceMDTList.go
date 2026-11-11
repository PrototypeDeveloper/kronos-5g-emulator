// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SliceMDTList represents ASN.1 SliceMDTList.
type SliceMDTList struct {
	List []SliceMDTItem `aper:"sizeLB:1,sizeUB:1024"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SliceMDTList) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("SliceMDTList", v) }

// DecodeSliceMDTList decodes a standalone APER value and reports the consumed octets.
func DecodeSliceMDTList(data []byte) (SliceMDTList, int, error) {
	var v SliceMDTList
	err := ngapcodec.DecodeParameter("SliceMDTList", data, &v)
	if ngapcodec.IsFatal(err) {
		return SliceMDTList{}, 0, err
	}
	return v, len(data), err
}
