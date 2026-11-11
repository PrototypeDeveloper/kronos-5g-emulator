// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExpectedUEMobility represents ASN.1 ExpectedUEMobility.
const (
	ExpectedUEMobilityPresentStationary Enumerated = 0
	ExpectedUEMobilityPresentMobile     Enumerated = 1
)

type ExpectedUEMobility struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExpectedUEMobility) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExpectedUEMobility", v)
}

// DecodeExpectedUEMobility decodes a standalone APER value and reports the consumed octets.
func DecodeExpectedUEMobility(data []byte) (ExpectedUEMobility, int, error) {
	var v ExpectedUEMobility
	err := ngapcodec.DecodeParameter("ExpectedUEMobility", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExpectedUEMobility{}, 0, err
	}
	return v, len(data), err
}
