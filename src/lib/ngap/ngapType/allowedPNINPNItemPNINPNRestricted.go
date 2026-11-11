// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AllowedPNINPNItemPNINPNRestricted represents ASN.1 AllowedPNINPNItemPNINPNRestricted.
const (
	AllowedPNINPNItemPNINPNRestrictedPresentRestricted    Enumerated = 0
	AllowedPNINPNItemPNINPNRestrictedPresentNotRestricted Enumerated = 1
)

type AllowedPNINPNItemPNINPNRestricted struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v AllowedPNINPNItemPNINPNRestricted) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:AllowedPNINPNItemPNINPNRestricted", v)
}

// DecodeAllowedPNINPNItemPNINPNRestricted decodes a standalone APER value and reports consumed octets.
func DecodeAllowedPNINPNItemPNINPNRestricted(data []byte) (AllowedPNINPNItemPNINPNRestricted, int, error) {
	var v AllowedPNINPNItemPNINPNRestricted
	err := ngapcodec.DecodeParameter("inline:AllowedPNINPNItemPNINPNRestricted", data, &v)
	if ngapcodec.IsFatal(err) {
		return AllowedPNINPNItemPNINPNRestricted{}, 0, err
	}
	return v, len(data), err
}
