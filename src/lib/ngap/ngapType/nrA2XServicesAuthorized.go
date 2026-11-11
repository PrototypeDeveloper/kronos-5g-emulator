// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRA2XServicesAuthorized represents ASN.1 NR-A2X-ServicesAuthorized.
type NRA2XServicesAuthorized struct {
	HasAerialUE           bool
	AerialUE              NRA2XServicesAuthorizedAerialUE `aper:"valueExt,valueLB:0,valueUB:1"`
	HasAerialControllerUE bool
	AerialControllerUE    NRA2XServicesAuthorizedAerialControllerUE `aper:"valueExt,valueLB:0,valueUB:1"`
	HasIEExtensions       bool
	IEExtensions          []NRA2XServicesAuthorizedExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRA2XServicesAuthorized) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NR-A2X-ServicesAuthorized", v)
}

// DecodeNRA2XServicesAuthorized decodes a standalone APER value and reports the consumed octets.
func DecodeNRA2XServicesAuthorized(data []byte) (NRA2XServicesAuthorized, int, error) {
	var v NRA2XServicesAuthorized
	err := ngapcodec.DecodeParameter("NR-A2X-ServicesAuthorized", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRA2XServicesAuthorized{}, 0, err
	}
	return v, len(data), err
}
