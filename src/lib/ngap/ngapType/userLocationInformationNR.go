// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UserLocationInformationNR represents ASN.1 UserLocationInformationNR.
type UserLocationInformationNR struct {
	NRCGI             NRCGI
	TAI               TAI
	HasTimeStamp      bool
	TimeStamp         TimeStamp
	HasIEExtensions   bool
	IEExtensions      []UserLocationInformationNRExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserLocationInformationNR) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserLocationInformationNR", v)
}

// DecodeUserLocationInformationNR decodes a standalone APER value and reports the consumed octets.
func DecodeUserLocationInformationNR(data []byte) (UserLocationInformationNR, int, error) {
	var v UserLocationInformationNR
	err := ngapcodec.DecodeParameter("UserLocationInformationNR", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserLocationInformationNR{}, 0, err
	}
	return v, len(data), err
}
