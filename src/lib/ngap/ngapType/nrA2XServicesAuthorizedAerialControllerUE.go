// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRA2XServicesAuthorizedAerialControllerUE represents ASN.1 NRA2XServicesAuthorizedAerialControllerUE.
const (
	NRA2XServicesAuthorizedAerialControllerUEPresentAuthorized    Enumerated = 0
	NRA2XServicesAuthorizedAerialControllerUEPresentNotAuthorized Enumerated = 1
)

type NRA2XServicesAuthorizedAerialControllerUE struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v NRA2XServicesAuthorizedAerialControllerUE) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:NRA2XServicesAuthorizedAerialControllerUE", v)
}

// DecodeNRA2XServicesAuthorizedAerialControllerUE decodes a standalone APER value and reports consumed octets.
func DecodeNRA2XServicesAuthorizedAerialControllerUE(data []byte) (NRA2XServicesAuthorizedAerialControllerUE, int, error) {
	var v NRA2XServicesAuthorizedAerialControllerUE
	err := ngapcodec.DecodeParameter("inline:NRA2XServicesAuthorizedAerialControllerUE", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRA2XServicesAuthorizedAerialControllerUE{}, 0, err
	}
	return v, len(data), err
}
