// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntendedServiceAreaCoordinates represents ASN.1 IntendedServiceAreaCoordinates.
type IntendedServiceAreaCoordinates struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntendedServiceAreaCoordinates) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntendedServiceAreaCoordinates", v)
}

// DecodeIntendedServiceAreaCoordinates decodes a standalone APER value and reports the consumed octets.
func DecodeIntendedServiceAreaCoordinates(data []byte) (IntendedServiceAreaCoordinates, int, error) {
	var v IntendedServiceAreaCoordinates
	err := ngapcodec.DecodeParameter("IntendedServiceAreaCoordinates", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntendedServiceAreaCoordinates{}, 0, err
	}
	return v, len(data), err
}
