// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UserLocationInformationEUTRA represents ASN.1 UserLocationInformationEUTRA.
type UserLocationInformationEUTRA struct {
	EUTRACGI          EUTRACGI
	TAI               TAI
	HasTimeStamp      bool
	TimeStamp         TimeStamp
	HasIEExtensions   bool
	IEExtensions      []UserLocationInformationEUTRAExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserLocationInformationEUTRA) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserLocationInformationEUTRA", v)
}

// DecodeUserLocationInformationEUTRA decodes a standalone APER value and reports the consumed octets.
func DecodeUserLocationInformationEUTRA(data []byte) (UserLocationInformationEUTRA, int, error) {
	var v UserLocationInformationEUTRA
	err := ngapcodec.DecodeParameter("UserLocationInformationEUTRA", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserLocationInformationEUTRA{}, 0, err
	}
	return v, len(data), err
}
