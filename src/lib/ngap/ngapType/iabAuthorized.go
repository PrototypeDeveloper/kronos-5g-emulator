// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IABAuthorized represents ASN.1 IAB-Authorized.
const (
	IABAuthorizedPresentAuthorized    Enumerated = 0
	IABAuthorizedPresentNotAuthorized Enumerated = 1
)

type IABAuthorized struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IABAuthorized) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IAB-Authorized", v)
}

// DecodeIABAuthorized decodes a standalone APER value and reports the consumed octets.
func DecodeIABAuthorized(data []byte) (IABAuthorized, int, error) {
	var v IABAuthorized
	err := ngapcodec.DecodeParameter("IAB-Authorized", data, &v)
	if ngapcodec.IsFatal(err) {
		return IABAuthorized{}, 0, err
	}
	return v, len(data), err
}
