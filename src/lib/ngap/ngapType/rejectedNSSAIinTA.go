// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RejectedNSSAIinTA represents ASN.1 RejectedNSSAIinTA.
type RejectedNSSAIinTA struct {
	Value []byte `aper:"sizeLB:32,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RejectedNSSAIinTA) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RejectedNSSAIinTA", v)
}

// DecodeRejectedNSSAIinTA decodes a standalone APER value and reports the consumed octets.
func DecodeRejectedNSSAIinTA(data []byte) (RejectedNSSAIinTA, int, error) {
	var v RejectedNSSAIinTA
	err := ngapcodec.DecodeParameter("RejectedNSSAIinTA", data, &v)
	if ngapcodec.IsFatal(err) {
		return RejectedNSSAIinTA{}, 0, err
	}
	return v, len(data), err
}
