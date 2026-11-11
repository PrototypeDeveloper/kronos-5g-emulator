// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIoTNASPDU represents ASN.1 AIoT-NASPDU.
type AIoTNASPDU struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTNASPDU) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("AIoT-NASPDU", v) }

// DecodeAIoTNASPDU decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTNASPDU(data []byte) (AIoTNASPDU, int, error) {
	var v AIoTNASPDU
	err := ngapcodec.DecodeParameter("AIoT-NASPDU", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTNASPDU{}, 0, err
	}
	return v, len(data), err
}
