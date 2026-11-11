// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MeasurementsToActivate represents ASN.1 MeasurementsToActivate.
type MeasurementsToActivate struct {
	Value BitString `aper:"sizeLB:8,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MeasurementsToActivate) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MeasurementsToActivate", v)
}

// DecodeMeasurementsToActivate decodes a standalone APER value and reports the consumed octets.
func DecodeMeasurementsToActivate(data []byte) (MeasurementsToActivate, int, error) {
	var v MeasurementsToActivate
	err := ngapcodec.DecodeParameter("MeasurementsToActivate", data, &v)
	if ngapcodec.IsFatal(err) {
		return MeasurementsToActivate{}, 0, err
	}
	return v, len(data), err
}
