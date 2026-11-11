// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RedirectionVoiceFallback represents ASN.1 RedirectionVoiceFallback.
const (
	RedirectionVoiceFallbackPresentPossible    Enumerated = 0
	RedirectionVoiceFallbackPresentNotPossible Enumerated = 1
)

type RedirectionVoiceFallback struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RedirectionVoiceFallback) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RedirectionVoiceFallback", v)
}

// DecodeRedirectionVoiceFallback decodes a standalone APER value and reports the consumed octets.
func DecodeRedirectionVoiceFallback(data []byte) (RedirectionVoiceFallback, int, error) {
	var v RedirectionVoiceFallback
	err := ngapcodec.DecodeParameter("RedirectionVoiceFallback", data, &v)
	if ngapcodec.IsFatal(err) {
		return RedirectionVoiceFallback{}, 0, err
	}
	return v, len(data), err
}
