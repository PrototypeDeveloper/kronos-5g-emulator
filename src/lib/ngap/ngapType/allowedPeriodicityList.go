// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AllowedPeriodicityList represents ASN.1 AllowedPeriodicityList.
type AllowedPeriodicityList struct {
	List []Periodicity `aper:"sizeLB:1,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AllowedPeriodicityList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AllowedPeriodicityList", v)
}

// DecodeAllowedPeriodicityList decodes a standalone APER value and reports the consumed octets.
func DecodeAllowedPeriodicityList(data []byte) (AllowedPeriodicityList, int, error) {
	var v AllowedPeriodicityList
	err := ngapcodec.DecodeParameter("AllowedPeriodicityList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AllowedPeriodicityList{}, 0, err
	}
	return v, len(data), err
}
