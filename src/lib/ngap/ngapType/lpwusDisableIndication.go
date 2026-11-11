// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LPWUSDisableIndication represents ASN.1 LPWUSDisableIndication.
const (
	LPWUSDisableIndicationPresentTrue Enumerated = 0
)

type LPWUSDisableIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LPWUSDisableIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LPWUSDisableIndication", v)
}

// DecodeLPWUSDisableIndication decodes a standalone APER value and reports the consumed octets.
func DecodeLPWUSDisableIndication(data []byte) (LPWUSDisableIndication, int, error) {
	var v LPWUSDisableIndication
	err := ngapcodec.DecodeParameter("LPWUSDisableIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return LPWUSDisableIndication{}, 0, err
	}
	return v, len(data), err
}
