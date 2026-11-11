// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CEmodeBrestricted represents ASN.1 CEmodeBrestricted.
const (
	CEmodeBrestrictedPresentRestricted    Enumerated = 0
	CEmodeBrestrictedPresentNotRestricted Enumerated = 1
)

type CEmodeBrestricted struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CEmodeBrestricted) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CEmodeBrestricted", v)
}

// DecodeCEmodeBrestricted decodes a standalone APER value and reports the consumed octets.
func DecodeCEmodeBrestricted(data []byte) (CEmodeBrestricted, int, error) {
	var v CEmodeBrestricted
	err := ngapcodec.DecodeParameter("CEmodeBrestricted", data, &v)
	if ngapcodec.IsFatal(err) {
		return CEmodeBrestricted{}, 0, err
	}
	return v, len(data), err
}
