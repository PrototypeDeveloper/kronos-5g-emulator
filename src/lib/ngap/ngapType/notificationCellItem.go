// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NotificationCellItem represents ASN.1 NotificationCell-Item.
type NotificationCellItem struct {
	NGRANCGI          NGRANCGI
	NotifyFlag        NotificationCellItemNotifyFlag `aper:"valueExt,valueLB:0,valueUB:1"`
	HasIEExtensions   bool
	IEExtensions      []NotificationCellItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NotificationCellItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NotificationCell-Item", v)
}

// DecodeNotificationCellItem decodes a standalone APER value and reports the consumed octets.
func DecodeNotificationCellItem(data []byte) (NotificationCellItem, int, error) {
	var v NotificationCellItem
	err := ngapcodec.DecodeParameter("NotificationCell-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return NotificationCellItem{}, 0, err
	}
	return v, len(data), err
}
