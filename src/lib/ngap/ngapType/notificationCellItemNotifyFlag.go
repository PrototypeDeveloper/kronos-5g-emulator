// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NotificationCellItemNotifyFlag represents ASN.1 NotificationCellItemNotifyFlag.
const (
	NotificationCellItemNotifyFlagPresentActivated   Enumerated = 0
	NotificationCellItemNotifyFlagPresentDeactivated Enumerated = 1
)

type NotificationCellItemNotifyFlag struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v NotificationCellItemNotifyFlag) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:NotificationCellItemNotifyFlag", v)
}

// DecodeNotificationCellItemNotifyFlag decodes a standalone APER value and reports consumed octets.
func DecodeNotificationCellItemNotifyFlag(data []byte) (NotificationCellItemNotifyFlag, int, error) {
	var v NotificationCellItemNotifyFlag
	err := ngapcodec.DecodeParameter("inline:NotificationCellItemNotifyFlag", data, &v)
	if ngapcodec.IsFatal(err) {
		return NotificationCellItemNotifyFlag{}, 0, err
	}
	return v, len(data), err
}
