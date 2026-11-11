// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRA2XServicesAuthorizedAerialUE represents ASN.1 NRA2XServicesAuthorizedAerialUE.
const (
	NRA2XServicesAuthorizedAerialUEPresentAuthorized    Enumerated = 0
	NRA2XServicesAuthorizedAerialUEPresentNotAuthorized Enumerated = 1
)

type NRA2XServicesAuthorizedAerialUE struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v NRA2XServicesAuthorizedAerialUE) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:NRA2XServicesAuthorizedAerialUE", v)
}

// DecodeNRA2XServicesAuthorizedAerialUE decodes a standalone APER value and reports consumed octets.
func DecodeNRA2XServicesAuthorizedAerialUE(data []byte) (NRA2XServicesAuthorizedAerialUE, int, error) {
	var v NRA2XServicesAuthorizedAerialUE
	err := ngapcodec.DecodeParameter("inline:NRA2XServicesAuthorizedAerialUE", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRA2XServicesAuthorizedAerialUE{}, 0, err
	}
	return v, len(data), err
}
