// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// WLANMeasConfig represents ASN.1 WLANMeasConfig.
const (
	WLANMeasConfigPresentSetup Enumerated = 0
)

type WLANMeasConfig struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v WLANMeasConfig) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("WLANMeasConfig", v)
}

// DecodeWLANMeasConfig decodes a standalone APER value and reports the consumed octets.
func DecodeWLANMeasConfig(data []byte) (WLANMeasConfig, int, error) {
	var v WLANMeasConfig
	err := ngapcodec.DecodeParameter("WLANMeasConfig", data, &v)
	if ngapcodec.IsFatal(err) {
		return WLANMeasConfig{}, 0, err
	}
	return v, len(data), err
}
