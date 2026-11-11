// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DAPSRequestInfoDAPSIndicator represents ASN.1 DAPSRequestInfoDAPSIndicator.
const (
	DAPSRequestInfoDAPSIndicatorPresentDapsHoRequired Enumerated = 0
)

type DAPSRequestInfoDAPSIndicator struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v DAPSRequestInfoDAPSIndicator) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:DAPSRequestInfoDAPSIndicator", v)
}

// DecodeDAPSRequestInfoDAPSIndicator decodes a standalone APER value and reports consumed octets.
func DecodeDAPSRequestInfoDAPSIndicator(data []byte) (DAPSRequestInfoDAPSIndicator, int, error) {
	var v DAPSRequestInfoDAPSIndicator
	err := ngapcodec.DecodeParameter("inline:DAPSRequestInfoDAPSIndicator", data, &v)
	if ngapcodec.IsFatal(err) {
		return DAPSRequestInfoDAPSIndicator{}, 0, err
	}
	return v, len(data), err
}
