// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// WarningAreaCoordinates represents ASN.1 WarningAreaCoordinates.
type WarningAreaCoordinates struct {
	Value []byte `aper:"sizeLB:1,sizeUB:1024"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v WarningAreaCoordinates) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("WarningAreaCoordinates", v)
}

// DecodeWarningAreaCoordinates decodes a standalone APER value and reports the consumed octets.
func DecodeWarningAreaCoordinates(data []byte) (WarningAreaCoordinates, int, error) {
	var v WarningAreaCoordinates
	err := ngapcodec.DecodeParameter("WarningAreaCoordinates", data, &v)
	if ngapcodec.IsFatal(err) {
		return WarningAreaCoordinates{}, 0, err
	}
	return v, len(data), err
}
