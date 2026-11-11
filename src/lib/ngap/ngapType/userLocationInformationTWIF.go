// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UserLocationInformationTWIF represents ASN.1 UserLocationInformationTWIF.
type UserLocationInformationTWIF struct {
	TWAPID            TWAPID
	IPAddress         TransportLayerAddress
	HasPortNumber     bool
	PortNumber        PortNumber
	HasIEExtensions   bool
	IEExtensions      []UserLocationInformationTWIFExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserLocationInformationTWIF) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserLocationInformationTWIF", v)
}

// DecodeUserLocationInformationTWIF decodes a standalone APER value and reports the consumed octets.
func DecodeUserLocationInformationTWIF(data []byte) (UserLocationInformationTWIF, int, error) {
	var v UserLocationInformationTWIF
	err := ngapcodec.DecodeParameter("UserLocationInformationTWIF", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserLocationInformationTWIF{}, 0, err
	}
	return v, len(data), err
}
