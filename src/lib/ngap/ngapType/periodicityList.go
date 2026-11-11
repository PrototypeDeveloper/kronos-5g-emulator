// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PeriodicityList represents ASN.1 PeriodicityList.
type PeriodicityList struct {
	AllowedPeriodicityList AllowedPeriodicityList
	HasIEExtensions        bool
	IEExtensions           []PeriodicityListExtIEsEntry
	UnknownExtensions      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PeriodicityList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PeriodicityList", v)
}

// DecodePeriodicityList decodes a standalone APER value and reports the consumed octets.
func DecodePeriodicityList(data []byte) (PeriodicityList, int, error) {
	var v PeriodicityList
	err := ngapcodec.DecodeParameter("PeriodicityList", data, &v)
	if ngapcodec.IsFatal(err) {
		return PeriodicityList{}, 0, err
	}
	return v, len(data), err
}
