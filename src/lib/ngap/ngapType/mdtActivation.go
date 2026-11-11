// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MDTActivation represents ASN.1 MDT-Activation.
const (
	MDTActivationPresentImmediateMDTOnly     Enumerated = 0
	MDTActivationPresentLoggedMDTOnly        Enumerated = 1
	MDTActivationPresentImmediateMDTAndTrace Enumerated = 2
)

type MDTActivation struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTActivation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MDT-Activation", v)
}

// DecodeMDTActivation decodes a standalone APER value and reports the consumed octets.
func DecodeMDTActivation(data []byte) (MDTActivation, int, error) {
	var v MDTActivation
	err := ngapcodec.DecodeParameter("MDT-Activation", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTActivation{}, 0, err
	}
	return v, len(data), err
}
