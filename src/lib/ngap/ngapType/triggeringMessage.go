// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TriggeringMessage represents ASN.1 TriggeringMessage.
const (
	TriggeringMessagePresentInitiatingMessage   Enumerated = 0
	TriggeringMessagePresentSuccessfulOutcome   Enumerated = 1
	TriggeringMessagePresentUnsuccessfulOutcome Enumerated = 2
)

type TriggeringMessage struct {
	Value Enumerated `aper:"valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TriggeringMessage) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TriggeringMessage", v)
}

// DecodeTriggeringMessage decodes a standalone APER value and reports the consumed octets.
func DecodeTriggeringMessage(data []byte) (TriggeringMessage, int, error) {
	var v TriggeringMessage
	err := ngapcodec.DecodeParameter("TriggeringMessage", data, &v)
	if ngapcodec.IsFatal(err) {
		return TriggeringMessage{}, 0, err
	}
	return v, len(data), err
}
