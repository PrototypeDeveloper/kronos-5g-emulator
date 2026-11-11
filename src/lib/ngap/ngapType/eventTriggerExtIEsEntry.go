// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EventTriggerExtIEsEntryPresent uint32

const (
	EventTriggerExtIEsEntryPresentNothing EventTriggerExtIEsEntryPresent = iota
	EventTriggerExtIEsEntryPresentUnknown
)

type EventTriggerExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      EventTriggerExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EventTriggerExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EventTrigger-ExtIEsEntry", v)
}

// DecodeEventTriggerExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEventTriggerExtIEsEntry(data []byte) (EventTriggerExtIEsEntry, int, error) {
	var v EventTriggerExtIEsEntry
	err := ngapcodec.DecodeParameter("EventTrigger-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EventTriggerExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
