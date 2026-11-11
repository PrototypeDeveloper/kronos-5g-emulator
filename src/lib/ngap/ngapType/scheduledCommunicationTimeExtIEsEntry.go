// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ScheduledCommunicationTimeExtIEsEntryPresent uint32

const (
	ScheduledCommunicationTimeExtIEsEntryPresentNothing ScheduledCommunicationTimeExtIEsEntryPresent = iota
	ScheduledCommunicationTimeExtIEsEntryPresentUnknown
)

type ScheduledCommunicationTimeExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      ScheduledCommunicationTimeExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ScheduledCommunicationTimeExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ScheduledCommunicationTime-ExtIEsEntry", v)
}

// DecodeScheduledCommunicationTimeExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeScheduledCommunicationTimeExtIEsEntry(data []byte) (ScheduledCommunicationTimeExtIEsEntry, int, error) {
	var v ScheduledCommunicationTimeExtIEsEntry
	err := ngapcodec.DecodeParameter("ScheduledCommunicationTime-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ScheduledCommunicationTimeExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
