// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UserPlaneSecurityInformation represents ASN.1 UserPlaneSecurityInformation.
type UserPlaneSecurityInformation struct {
	SecurityResult     SecurityResult
	SecurityIndication SecurityIndication
	HasIEExtensions    bool
	IEExtensions       []UserPlaneSecurityInformationExtIEsEntry
	UnknownExtensions  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserPlaneSecurityInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserPlaneSecurityInformation", v)
}

// DecodeUserPlaneSecurityInformation decodes a standalone APER value and reports the consumed octets.
func DecodeUserPlaneSecurityInformation(data []byte) (UserPlaneSecurityInformation, int, error) {
	var v UserPlaneSecurityInformation
	err := ngapcodec.DecodeParameter("UserPlaneSecurityInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserPlaneSecurityInformation{}, 0, err
	}
	return v, len(data), err
}
