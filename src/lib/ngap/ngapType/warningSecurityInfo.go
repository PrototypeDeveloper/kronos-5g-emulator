// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// WarningSecurityInfo represents ASN.1 WarningSecurityInfo.
type WarningSecurityInfo struct {
	Value []byte `aper:"sizeLB:50,sizeUB:50"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v WarningSecurityInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("WarningSecurityInfo", v)
}

// DecodeWarningSecurityInfo decodes a standalone APER value and reports the consumed octets.
func DecodeWarningSecurityInfo(data []byte) (WarningSecurityInfo, int, error) {
	var v WarningSecurityInfo
	err := ngapcodec.DecodeParameter("WarningSecurityInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return WarningSecurityInfo{}, 0, err
	}
	return v, len(data), err
}
