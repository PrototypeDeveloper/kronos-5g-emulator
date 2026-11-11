// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NotificationCellList represents ASN.1 NotificationCellList.
type NotificationCellList struct {
	List []NotificationCellItem `aper:"sizeLB:1,sizeUB:16384"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NotificationCellList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NotificationCellList", v)
}

// DecodeNotificationCellList decodes a standalone APER value and reports the consumed octets.
func DecodeNotificationCellList(data []byte) (NotificationCellList, int, error) {
	var v NotificationCellList
	err := ngapcodec.DecodeParameter("NotificationCellList", data, &v)
	if ngapcodec.IsFatal(err) {
		return NotificationCellList{}, 0, err
	}
	return v, len(data), err
}
