// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// VehicleUE represents ASN.1 VehicleUE.
const (
	VehicleUEPresentAuthorized    Enumerated = 0
	VehicleUEPresentNotAuthorized Enumerated = 1
)

type VehicleUE struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v VehicleUE) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("VehicleUE", v) }

// DecodeVehicleUE decodes a standalone APER value and reports the consumed octets.
func DecodeVehicleUE(data []byte) (VehicleUE, int, error) {
	var v VehicleUE
	err := ngapcodec.DecodeParameter("VehicleUE", data, &v)
	if ngapcodec.IsFatal(err) {
		return VehicleUE{}, 0, err
	}
	return v, len(data), err
}
