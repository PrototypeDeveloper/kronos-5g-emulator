// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RejectedNSSAIinPLMN represents ASN.1 RejectedNSSAIinPLMN.
type RejectedNSSAIinPLMN struct {
	Value []byte `aper:"sizeLB:32,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RejectedNSSAIinPLMN) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RejectedNSSAIinPLMN", v)
}

// DecodeRejectedNSSAIinPLMN decodes a standalone APER value and reports the consumed octets.
func DecodeRejectedNSSAIinPLMN(data []byte) (RejectedNSSAIinPLMN, int, error) {
	var v RejectedNSSAIinPLMN
	err := ngapcodec.DecodeParameter("RejectedNSSAIinPLMN", data, &v)
	if ngapcodec.IsFatal(err) {
		return RejectedNSSAIinPLMN{}, 0, err
	}
	return v, len(data), err
}
