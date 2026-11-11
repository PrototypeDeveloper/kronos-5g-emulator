// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MobileIABSupported represents ASN.1 MobileIAB-Supported.
const (
	MobileIABSupportedPresentTrue Enumerated = 0
)

type MobileIABSupported struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MobileIABSupported) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MobileIAB-Supported", v)
}

// DecodeMobileIABSupported decodes a standalone APER value and reports the consumed octets.
func DecodeMobileIABSupported(data []byte) (MobileIABSupported, int, error) {
	var v MobileIABSupported
	err := ngapcodec.DecodeParameter("MobileIAB-Supported", data, &v)
	if ngapcodec.IsFatal(err) {
		return MobileIABSupported{}, 0, err
	}
	return v, len(data), err
}
