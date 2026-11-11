// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGAPIESupportInformationRequestList represents ASN.1 NGAPIESupportInformationRequestList.
type NGAPIESupportInformationRequestList struct {
	List []NGAPIESupportInformationRequestItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGAPIESupportInformationRequestList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGAPIESupportInformationRequestList", v)
}

// DecodeNGAPIESupportInformationRequestList decodes a standalone APER value and reports the consumed octets.
func DecodeNGAPIESupportInformationRequestList(data []byte) (NGAPIESupportInformationRequestList, int, error) {
	var v NGAPIESupportInformationRequestList
	err := ngapcodec.DecodeParameter("NGAPIESupportInformationRequestList", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGAPIESupportInformationRequestList{}, 0, err
	}
	return v, len(data), err
}
