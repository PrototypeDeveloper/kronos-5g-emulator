// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UserLocationInformationTNGF represents ASN.1 UserLocationInformationTNGF.
type UserLocationInformationTNGF struct {
	TNAPID            TNAPID
	IPAddress         TransportLayerAddress
	HasPortNumber     bool
	PortNumber        PortNumber
	HasIEExtensions   bool
	IEExtensions      []UserLocationInformationTNGFExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserLocationInformationTNGF) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserLocationInformationTNGF", v)
}

// DecodeUserLocationInformationTNGF decodes a standalone APER value and reports the consumed octets.
func DecodeUserLocationInformationTNGF(data []byte) (UserLocationInformationTNGF, int, error) {
	var v UserLocationInformationTNGF
	err := ngapcodec.DecodeParameter("UserLocationInformationTNGF", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserLocationInformationTNGF{}, 0, err
	}
	return v, len(data), err
}
