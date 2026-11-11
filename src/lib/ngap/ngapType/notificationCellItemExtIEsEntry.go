// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NotificationCellItemExtIEsEntryPresent uint32

const (
	NotificationCellItemExtIEsEntryPresentNothing NotificationCellItemExtIEsEntryPresent = iota
	NotificationCellItemExtIEsEntryPresentUnknown
)

type NotificationCellItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NotificationCellItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NotificationCellItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NotificationCell-Item-ExtIEsEntry", v)
}

// DecodeNotificationCellItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNotificationCellItemExtIEsEntry(data []byte) (NotificationCellItemExtIEsEntry, int, error) {
	var v NotificationCellItemExtIEsEntry
	err := ngapcodec.DecodeParameter("NotificationCell-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NotificationCellItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
