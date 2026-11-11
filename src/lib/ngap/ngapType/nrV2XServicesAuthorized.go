// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRV2XServicesAuthorized represents ASN.1 NRV2XServicesAuthorized.
type NRV2XServicesAuthorized struct {
	HasVehicleUE      bool
	VehicleUE         VehicleUE
	HasPedestrianUE   bool
	PedestrianUE      PedestrianUE
	HasIEExtensions   bool
	IEExtensions      []NRV2XServicesAuthorizedExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRV2XServicesAuthorized) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NRV2XServicesAuthorized", v)
}

// DecodeNRV2XServicesAuthorized decodes a standalone APER value and reports the consumed octets.
func DecodeNRV2XServicesAuthorized(data []byte) (NRV2XServicesAuthorized, int, error) {
	var v NRV2XServicesAuthorized
	err := ngapcodec.DecodeParameter("NRV2XServicesAuthorized", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRV2XServicesAuthorized{}, 0, err
	}
	return v, len(data), err
}
