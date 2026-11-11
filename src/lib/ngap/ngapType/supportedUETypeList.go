// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SupportedUETypeList represents ASN.1 SupportedUETypeList.
type SupportedUETypeList struct {
	List []SupportedUEType `aper:"sizeLB:1,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SupportedUETypeList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SupportedUETypeList", v)
}

// DecodeSupportedUETypeList decodes a standalone APER value and reports the consumed octets.
func DecodeSupportedUETypeList(data []byte) (SupportedUETypeList, int, error) {
	var v SupportedUETypeList
	err := ngapcodec.DecodeParameter("SupportedUETypeList", data, &v)
	if ngapcodec.IsFatal(err) {
		return SupportedUETypeList{}, 0, err
	}
	return v, len(data), err
}
