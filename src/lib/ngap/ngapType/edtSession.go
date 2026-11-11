// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EDTSession represents ASN.1 EDT-Session.
const (
	EDTSessionPresentTrue Enumerated = 0
)

type EDTSession struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EDTSession) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("EDT-Session", v) }

// DecodeEDTSession decodes a standalone APER value and reports the consumed octets.
func DecodeEDTSession(data []byte) (EDTSession, int, error) {
	var v EDTSession
	err := ngapcodec.DecodeParameter("EDT-Session", data, &v)
	if ngapcodec.IsFatal(err) {
		return EDTSession{}, 0, err
	}
	return v, len(data), err
}
