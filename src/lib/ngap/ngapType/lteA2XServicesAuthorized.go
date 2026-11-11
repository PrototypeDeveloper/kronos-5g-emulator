// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LTEA2XServicesAuthorized represents ASN.1 LTE-A2X-ServicesAuthorized.
type LTEA2XServicesAuthorized struct {
	HasAerialUE           bool
	AerialUE              LTEA2XServicesAuthorizedAerialUE `aper:"valueExt,valueLB:0,valueUB:1"`
	HasAerialControllerUE bool
	AerialControllerUE    LTEA2XServicesAuthorizedAerialControllerUE `aper:"valueExt,valueLB:0,valueUB:1"`
	HasIEExtensions       bool
	IEExtensions          []LTEA2XServicesAuthorizedExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LTEA2XServicesAuthorized) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LTE-A2X-ServicesAuthorized", v)
}

// DecodeLTEA2XServicesAuthorized decodes a standalone APER value and reports the consumed octets.
func DecodeLTEA2XServicesAuthorized(data []byte) (LTEA2XServicesAuthorized, int, error) {
	var v LTEA2XServicesAuthorized
	err := ngapcodec.DecodeParameter("LTE-A2X-ServicesAuthorized", data, &v)
	if ngapcodec.IsFatal(err) {
		return LTEA2XServicesAuthorized{}, 0, err
	}
	return v, len(data), err
}
