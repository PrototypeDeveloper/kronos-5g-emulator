// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IncludeBeamMeasurementsIndication represents ASN.1 IncludeBeamMeasurementsIndication.
const (
	IncludeBeamMeasurementsIndicationPresentTrue Enumerated = 0
)

type IncludeBeamMeasurementsIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IncludeBeamMeasurementsIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IncludeBeamMeasurementsIndication", v)
}

// DecodeIncludeBeamMeasurementsIndication decodes a standalone APER value and reports the consumed octets.
func DecodeIncludeBeamMeasurementsIndication(data []byte) (IncludeBeamMeasurementsIndication, int, error) {
	var v IncludeBeamMeasurementsIndication
	err := ngapcodec.DecodeParameter("IncludeBeamMeasurementsIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return IncludeBeamMeasurementsIndication{}, 0, err
	}
	return v, len(data), err
}
