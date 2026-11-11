// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AerialUEsubscriptionInformation represents ASN.1 AerialUEsubscriptionInformation.
const (
	AerialUEsubscriptionInformationPresentAllowed    Enumerated = 0
	AerialUEsubscriptionInformationPresentNotAllowed Enumerated = 1
)

type AerialUEsubscriptionInformation struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AerialUEsubscriptionInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AerialUEsubscriptionInformation", v)
}

// DecodeAerialUEsubscriptionInformation decodes a standalone APER value and reports the consumed octets.
func DecodeAerialUEsubscriptionInformation(data []byte) (AerialUEsubscriptionInformation, int, error) {
	var v AerialUEsubscriptionInformation
	err := ngapcodec.DecodeParameter("AerialUEsubscriptionInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return AerialUEsubscriptionInformation{}, 0, err
	}
	return v, len(data), err
}
