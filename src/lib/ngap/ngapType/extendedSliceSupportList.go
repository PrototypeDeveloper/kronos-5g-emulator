// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExtendedSliceSupportList represents ASN.1 ExtendedSliceSupportList.
type ExtendedSliceSupportList struct {
	List []SliceSupportItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExtendedSliceSupportList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExtendedSliceSupportList", v)
}

// DecodeExtendedSliceSupportList decodes a standalone APER value and reports the consumed octets.
func DecodeExtendedSliceSupportList(data []byte) (ExtendedSliceSupportList, int, error) {
	var v ExtendedSliceSupportList
	err := ngapcodec.DecodeParameter("ExtendedSliceSupportList", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExtendedSliceSupportList{}, 0, err
	}
	return v, len(data), err
}
