// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGAPIESupportInformationResponseList represents ASN.1 NGAPIESupportInformationResponseList.
type NGAPIESupportInformationResponseList struct {
	List []NGAPIESupportInformationResponseItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGAPIESupportInformationResponseList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGAPIESupportInformationResponseList", v)
}

// DecodeNGAPIESupportInformationResponseList decodes a standalone APER value and reports the consumed octets.
func DecodeNGAPIESupportInformationResponseList(data []byte) (NGAPIESupportInformationResponseList, int, error) {
	var v NGAPIESupportInformationResponseList
	err := ngapcodec.DecodeParameter("NGAPIESupportInformationResponseList", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGAPIESupportInformationResponseList{}, 0, err
	}
	return v, len(data), err
}
