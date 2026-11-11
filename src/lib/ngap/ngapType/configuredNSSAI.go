// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ConfiguredNSSAI represents ASN.1 ConfiguredNSSAI.
type ConfiguredNSSAI struct {
	Value []byte `aper:"sizeLB:128,sizeUB:128"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ConfiguredNSSAI) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ConfiguredNSSAI", v)
}

// DecodeConfiguredNSSAI decodes a standalone APER value and reports the consumed octets.
func DecodeConfiguredNSSAI(data []byte) (ConfiguredNSSAI, int, error) {
	var v ConfiguredNSSAI
	err := ngapcodec.DecodeParameter("ConfiguredNSSAI", data, &v)
	if ngapcodec.IsFatal(err) {
		return ConfiguredNSSAI{}, 0, err
	}
	return v, len(data), err
}
