// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIoTCorrelationIdentifier represents ASN.1 AIoT-CorrelationIdentifier.
type AIoTCorrelationIdentifier struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTCorrelationIdentifier) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-CorrelationIdentifier", v)
}

// DecodeAIoTCorrelationIdentifier decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTCorrelationIdentifier(data []byte) (AIoTCorrelationIdentifier, int, error) {
	var v AIoTCorrelationIdentifier
	err := ngapcodec.DecodeParameter("AIoT-CorrelationIdentifier", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTCorrelationIdentifier{}, 0, err
	}
	return v, len(data), err
}
