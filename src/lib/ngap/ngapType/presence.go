// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// Presence represents ASN.1 Presence.
const (
	PresencePresentOptional    Enumerated = 0
	PresencePresentConditional Enumerated = 1
	PresencePresentMandatory   Enumerated = 2
)

type Presence struct {
	Value Enumerated `aper:"valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v Presence) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("Presence", v) }

// DecodePresence decodes a standalone APER value and reports the consumed octets.
func DecodePresence(data []byte) (Presence, int, error) {
	var v Presence
	err := ngapcodec.DecodeParameter("Presence", data, &v)
	if ngapcodec.IsFatal(err) {
		return Presence{}, 0, err
	}
	return v, len(data), err
}
