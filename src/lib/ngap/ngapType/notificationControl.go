// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NotificationControl represents ASN.1 NotificationControl.
const (
	NotificationControlPresentNotificationRequested Enumerated = 0
)

type NotificationControl struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NotificationControl) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NotificationControl", v)
}

// DecodeNotificationControl decodes a standalone APER value and reports the consumed octets.
func DecodeNotificationControl(data []byte) (NotificationControl, int, error) {
	var v NotificationControl
	err := ngapcodec.DecodeParameter("NotificationControl", data, &v)
	if ngapcodec.IsFatal(err) {
		return NotificationControl{}, 0, err
	}
	return v, len(data), err
}
