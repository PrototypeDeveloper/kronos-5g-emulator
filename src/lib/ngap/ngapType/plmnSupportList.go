// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PLMNSupportList represents ASN.1 PLMNSupportList.
type PLMNSupportList struct {
	List []PLMNSupportItem `aper:"sizeLB:1,sizeUB:12"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PLMNSupportList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PLMNSupportList", v)
}

// DecodePLMNSupportList decodes a standalone APER value and reports the consumed octets.
func DecodePLMNSupportList(data []byte) (PLMNSupportList, int, error) {
	var v PLMNSupportList
	err := ngapcodec.DecodeParameter("PLMNSupportList", data, &v)
	if ngapcodec.IsFatal(err) {
		return PLMNSupportList{}, 0, err
	}
	return v, len(data), err
}
