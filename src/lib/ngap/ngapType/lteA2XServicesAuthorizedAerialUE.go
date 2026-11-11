// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LTEA2XServicesAuthorizedAerialUE represents ASN.1 LTEA2XServicesAuthorizedAerialUE.
const (
	LTEA2XServicesAuthorizedAerialUEPresentAuthorized    Enumerated = 0
	LTEA2XServicesAuthorizedAerialUEPresentNotAuthorized Enumerated = 1
)

type LTEA2XServicesAuthorizedAerialUE struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v LTEA2XServicesAuthorizedAerialUE) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:LTEA2XServicesAuthorizedAerialUE", v)
}

// DecodeLTEA2XServicesAuthorizedAerialUE decodes a standalone APER value and reports consumed octets.
func DecodeLTEA2XServicesAuthorizedAerialUE(data []byte) (LTEA2XServicesAuthorizedAerialUE, int, error) {
	var v LTEA2XServicesAuthorizedAerialUE
	err := ngapcodec.DecodeParameter("inline:LTEA2XServicesAuthorizedAerialUE", data, &v)
	if ngapcodec.IsFatal(err) {
		return LTEA2XServicesAuthorizedAerialUE{}, 0, err
	}
	return v, len(data), err
}
