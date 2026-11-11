// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ConfiguredTACIndication represents ASN.1 ConfiguredTACIndication.
const (
	ConfiguredTACIndicationPresentTrue Enumerated = 0
)

type ConfiguredTACIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ConfiguredTACIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ConfiguredTACIndication", v)
}

// DecodeConfiguredTACIndication decodes a standalone APER value and reports the consumed octets.
func DecodeConfiguredTACIndication(data []byte) (ConfiguredTACIndication, int, error) {
	var v ConfiguredTACIndication
	err := ngapcodec.DecodeParameter("ConfiguredTACIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return ConfiguredTACIndication{}, 0, err
	}
	return v, len(data), err
}
