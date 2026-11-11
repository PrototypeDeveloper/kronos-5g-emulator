// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MobileIABAuthorized represents ASN.1 MobileIAB-Authorized.
const (
	MobileIABAuthorizedPresentAuthorized    Enumerated = 0
	MobileIABAuthorizedPresentNotAuthorized Enumerated = 1
)

type MobileIABAuthorized struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MobileIABAuthorized) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MobileIAB-Authorized", v)
}

// DecodeMobileIABAuthorized decodes a standalone APER value and reports the consumed octets.
func DecodeMobileIABAuthorized(data []byte) (MobileIABAuthorized, int, error) {
	var v MobileIABAuthorized
	err := ngapcodec.DecodeParameter("MobileIAB-Authorized", data, &v)
	if ngapcodec.IsFatal(err) {
		return MobileIABAuthorized{}, 0, err
	}
	return v, len(data), err
}
