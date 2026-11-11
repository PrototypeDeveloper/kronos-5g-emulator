// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TargetNSSAI represents ASN.1 TargetNSSAI.
type TargetNSSAI struct {
	List []TargetNSSAIItem `aper:"sizeLB:1,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetNSSAI) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TargetNSSAI", v) }

// DecodeTargetNSSAI decodes a standalone APER value and reports the consumed octets.
func DecodeTargetNSSAI(data []byte) (TargetNSSAI, int, error) {
	var v TargetNSSAI
	err := ngapcodec.DecodeParameter("TargetNSSAI", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetNSSAI{}, 0, err
	}
	return v, len(data), err
}
