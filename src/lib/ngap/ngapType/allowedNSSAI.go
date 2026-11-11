// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AllowedNSSAI represents ASN.1 AllowedNSSAI.
type AllowedNSSAI struct {
	List []AllowedNSSAIItem `aper:"sizeLB:1,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AllowedNSSAI) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("AllowedNSSAI", v) }

// DecodeAllowedNSSAI decodes a standalone APER value and reports the consumed octets.
func DecodeAllowedNSSAI(data []byte) (AllowedNSSAI, int, error) {
	var v AllowedNSSAI
	err := ngapcodec.DecodeParameter("AllowedNSSAI", data, &v)
	if ngapcodec.IsFatal(err) {
		return AllowedNSSAI{}, 0, err
	}
	return v, len(data), err
}
