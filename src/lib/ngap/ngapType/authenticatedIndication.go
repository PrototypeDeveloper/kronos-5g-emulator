// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AuthenticatedIndication represents ASN.1 AuthenticatedIndication.
const (
	AuthenticatedIndicationPresentTrue Enumerated = 0
)

type AuthenticatedIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AuthenticatedIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AuthenticatedIndication", v)
}

// DecodeAuthenticatedIndication decodes a standalone APER value and reports the consumed octets.
func DecodeAuthenticatedIndication(data []byte) (AuthenticatedIndication, int, error) {
	var v AuthenticatedIndication
	err := ngapcodec.DecodeParameter("AuthenticatedIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return AuthenticatedIndication{}, 0, err
	}
	return v, len(data), err
}
