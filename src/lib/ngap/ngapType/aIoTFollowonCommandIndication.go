// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIoTFollowonCommandIndication represents ASN.1 AIoT-FollowonCommandIndication.
const (
	AIoTFollowonCommandIndicationPresentTrue Enumerated = 0
)

type AIoTFollowonCommandIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTFollowonCommandIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-FollowonCommandIndication", v)
}

// DecodeAIoTFollowonCommandIndication decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTFollowonCommandIndication(data []byte) (AIoTFollowonCommandIndication, int, error) {
	var v AIoTFollowonCommandIndication
	err := ngapcodec.DecodeParameter("AIoT-FollowonCommandIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTFollowonCommandIndication{}, 0, err
	}
	return v, len(data), err
}
