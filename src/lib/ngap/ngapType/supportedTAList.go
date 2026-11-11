// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SupportedTAList represents ASN.1 SupportedTAList.
type SupportedTAList struct {
	List []SupportedTAItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SupportedTAList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SupportedTAList", v)
}

// DecodeSupportedTAList decodes a standalone APER value and reports the consumed octets.
func DecodeSupportedTAList(data []byte) (SupportedTAList, int, error) {
	var v SupportedTAList
	err := ngapcodec.DecodeParameter("SupportedTAList", data, &v)
	if ngapcodec.IsFatal(err) {
		return SupportedTAList{}, 0, err
	}
	return v, len(data), err
}
