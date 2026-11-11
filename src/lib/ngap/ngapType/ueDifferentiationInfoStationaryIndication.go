// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEDifferentiationInfoStationaryIndication represents ASN.1 UEDifferentiationInfoStationaryIndication.
const (
	UEDifferentiationInfoStationaryIndicationPresentStationary Enumerated = 0
	UEDifferentiationInfoStationaryIndicationPresentMobile     Enumerated = 1
)

type UEDifferentiationInfoStationaryIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v UEDifferentiationInfoStationaryIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:UEDifferentiationInfoStationaryIndication", v)
}

// DecodeUEDifferentiationInfoStationaryIndication decodes a standalone APER value and reports consumed octets.
func DecodeUEDifferentiationInfoStationaryIndication(data []byte) (UEDifferentiationInfoStationaryIndication, int, error) {
	var v UEDifferentiationInfoStationaryIndication
	err := ngapcodec.DecodeParameter("inline:UEDifferentiationInfoStationaryIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEDifferentiationInfoStationaryIndication{}, 0, err
	}
	return v, len(data), err
}
