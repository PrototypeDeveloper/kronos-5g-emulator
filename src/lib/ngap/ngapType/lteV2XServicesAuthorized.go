// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LTEV2XServicesAuthorized represents ASN.1 LTEV2XServicesAuthorized.
type LTEV2XServicesAuthorized struct {
	HasVehicleUE      bool
	VehicleUE         VehicleUE
	HasPedestrianUE   bool
	PedestrianUE      PedestrianUE
	HasIEExtensions   bool
	IEExtensions      []LTEV2XServicesAuthorizedExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LTEV2XServicesAuthorized) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LTEV2XServicesAuthorized", v)
}

// DecodeLTEV2XServicesAuthorized decodes a standalone APER value and reports the consumed octets.
func DecodeLTEV2XServicesAuthorized(data []byte) (LTEV2XServicesAuthorized, int, error) {
	var v LTEV2XServicesAuthorized
	err := ngapcodec.DecodeParameter("LTEV2XServicesAuthorized", data, &v)
	if ngapcodec.IsFatal(err) {
		return LTEV2XServicesAuthorized{}, 0, err
	}
	return v, len(data), err
}
