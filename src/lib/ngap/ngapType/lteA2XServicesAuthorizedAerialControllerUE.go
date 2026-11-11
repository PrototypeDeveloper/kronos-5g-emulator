// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LTEA2XServicesAuthorizedAerialControllerUE represents ASN.1 LTEA2XServicesAuthorizedAerialControllerUE.
const (
	LTEA2XServicesAuthorizedAerialControllerUEPresentAuthorized    Enumerated = 0
	LTEA2XServicesAuthorizedAerialControllerUEPresentNotAuthorized Enumerated = 1
)

type LTEA2XServicesAuthorizedAerialControllerUE struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v LTEA2XServicesAuthorizedAerialControllerUE) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:LTEA2XServicesAuthorizedAerialControllerUE", v)
}

// DecodeLTEA2XServicesAuthorizedAerialControllerUE decodes a standalone APER value and reports consumed octets.
func DecodeLTEA2XServicesAuthorizedAerialControllerUE(data []byte) (LTEA2XServicesAuthorizedAerialControllerUE, int, error) {
	var v LTEA2XServicesAuthorizedAerialControllerUE
	err := ngapcodec.DecodeParameter("inline:LTEA2XServicesAuthorizedAerialControllerUE", data, &v)
	if ngapcodec.IsFatal(err) {
		return LTEA2XServicesAuthorizedAerialControllerUE{}, 0, err
	}
	return v, len(data), err
}
