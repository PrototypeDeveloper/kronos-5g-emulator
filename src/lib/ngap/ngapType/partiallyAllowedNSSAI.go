// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PartiallyAllowedNSSAI represents ASN.1 Partially-Allowed-NSSAI.
type PartiallyAllowedNSSAI struct {
	List []PartiallyAllowedNSSAIItem `aper:"sizeLB:1,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PartiallyAllowedNSSAI) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Partially-Allowed-NSSAI", v)
}

// DecodePartiallyAllowedNSSAI decodes a standalone APER value and reports the consumed octets.
func DecodePartiallyAllowedNSSAI(data []byte) (PartiallyAllowedNSSAI, int, error) {
	var v PartiallyAllowedNSSAI
	err := ngapcodec.DecodeParameter("Partially-Allowed-NSSAI", data, &v)
	if ngapcodec.IsFatal(err) {
		return PartiallyAllowedNSSAI{}, 0, err
	}
	return v, len(data), err
}
