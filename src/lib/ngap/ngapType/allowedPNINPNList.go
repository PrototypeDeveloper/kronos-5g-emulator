// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AllowedPNINPNList represents ASN.1 Allowed-PNI-NPN-List.
type AllowedPNINPNList struct {
	List []AllowedPNINPNItem `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AllowedPNINPNList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Allowed-PNI-NPN-List", v)
}

// DecodeAllowedPNINPNList decodes a standalone APER value and reports the consumed octets.
func DecodeAllowedPNINPNList(data []byte) (AllowedPNINPNList, int, error) {
	var v AllowedPNINPNList
	err := ngapcodec.DecodeParameter("Allowed-PNI-NPN-List", data, &v)
	if ngapcodec.IsFatal(err) {
		return AllowedPNINPNList{}, 0, err
	}
	return v, len(data), err
}
