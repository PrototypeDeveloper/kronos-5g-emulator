// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RequestedNSSAI represents ASN.1 RequestedNSSAI.
type RequestedNSSAI struct {
	List []RequestedNSSAIItem `aper:"sizeLB:1,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RequestedNSSAI) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RequestedNSSAI", v)
}

// DecodeRequestedNSSAI decodes a standalone APER value and reports the consumed octets.
func DecodeRequestedNSSAI(data []byte) (RequestedNSSAI, int, error) {
	var v RequestedNSSAI
	err := ngapcodec.DecodeParameter("RequestedNSSAI", data, &v)
	if ngapcodec.IsFatal(err) {
		return RequestedNSSAI{}, 0, err
	}
	return v, len(data), err
}
