// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NotificationCause represents ASN.1 NotificationCause.
const (
	NotificationCausePresentFulfilled      Enumerated = 0
	NotificationCausePresentNotFulfilled   Enumerated = 1
	NotificationCausePresentNotFulfilledDL Enumerated = 2
	NotificationCausePresentNotFulfilledUL Enumerated = 3
)

type NotificationCause struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NotificationCause) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NotificationCause", v)
}

// DecodeNotificationCause decodes a standalone APER value and reports the consumed octets.
func DecodeNotificationCause(data []byte) (NotificationCause, int, error) {
	var v NotificationCause
	err := ngapcodec.DecodeParameter("NotificationCause", data, &v)
	if ngapcodec.IsFatal(err) {
		return NotificationCause{}, 0, err
	}
	return v, len(data), err
}
