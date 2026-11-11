// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UserLocationInformationN3IWFWithoutPortNumber represents ASN.1 UserLocationInformationN3IWF-without-PortNumber.
type UserLocationInformationN3IWFWithoutPortNumber struct {
	IPAddress         TransportLayerAddress
	HasTAI            bool
	TAI               TAI
	HasIEExtensions   bool
	IEExtensions      []UserLocationInformationN3IWFWithoutPortNumberExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserLocationInformationN3IWFWithoutPortNumber) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserLocationInformationN3IWF-without-PortNumber", v)
}

// DecodeUserLocationInformationN3IWFWithoutPortNumber decodes a standalone APER value and reports the consumed octets.
func DecodeUserLocationInformationN3IWFWithoutPortNumber(data []byte) (UserLocationInformationN3IWFWithoutPortNumber, int, error) {
	var v UserLocationInformationN3IWFWithoutPortNumber
	err := ngapcodec.DecodeParameter("UserLocationInformationN3IWF-without-PortNumber", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserLocationInformationN3IWFWithoutPortNumber{}, 0, err
	}
	return v, len(data), err
}
