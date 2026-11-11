// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IABSupported represents ASN.1 IAB-Supported.
const (
	IABSupportedPresentTrue Enumerated = 0
)

type IABSupported struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IABSupported) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("IAB-Supported", v) }

// DecodeIABSupported decodes a standalone APER value and reports the consumed octets.
func DecodeIABSupported(data []byte) (IABSupported, int, error) {
	var v IABSupported
	err := ngapcodec.DecodeParameter("IAB-Supported", data, &v)
	if ngapcodec.IsFatal(err) {
		return IABSupported{}, 0, err
	}
	return v, len(data), err
}
