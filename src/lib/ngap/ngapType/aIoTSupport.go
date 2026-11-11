// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIoTSupport represents ASN.1 AIoT-Support.
const (
	AIoTSupportPresentAIoTOnly    Enumerated = 0
	AIoTSupportPresentAIoTAndNRUu Enumerated = 1
)

type AIoTSupport struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTSupport) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("AIoT-Support", v) }

// DecodeAIoTSupport decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTSupport(data []byte) (AIoTSupport, int, error) {
	var v AIoTSupport
	err := ngapcodec.DecodeParameter("AIoT-Support", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTSupport{}, 0, err
	}
	return v, len(data), err
}
