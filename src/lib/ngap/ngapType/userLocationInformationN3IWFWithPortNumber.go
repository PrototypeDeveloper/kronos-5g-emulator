// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UserLocationInformationN3IWFWithPortNumber represents ASN.1 UserLocationInformationN3IWF-with-PortNumber.
type UserLocationInformationN3IWFWithPortNumber struct {
	IPAddress         TransportLayerAddress
	PortNumber        PortNumber
	HasIEExtensions   bool
	IEExtensions      []UserLocationInformationN3IWFWithPortNumberExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserLocationInformationN3IWFWithPortNumber) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserLocationInformationN3IWF-with-PortNumber", v)
}

// DecodeUserLocationInformationN3IWFWithPortNumber decodes a standalone APER value and reports the consumed octets.
func DecodeUserLocationInformationN3IWFWithPortNumber(data []byte) (UserLocationInformationN3IWFWithPortNumber, int, error) {
	var v UserLocationInformationN3IWFWithPortNumber
	err := ngapcodec.DecodeParameter("UserLocationInformationN3IWF-with-PortNumber", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserLocationInformationN3IWFWithPortNumber{}, 0, err
	}
	return v, len(data), err
}
