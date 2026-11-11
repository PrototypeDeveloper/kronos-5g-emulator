// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EventBasedReportingIEsExtIEsEntryPresent uint32

const (
	EventBasedReportingIEsExtIEsEntryPresentNothing EventBasedReportingIEsExtIEsEntryPresent = iota
	EventBasedReportingIEsExtIEsEntryPresentUnknown
)

type EventBasedReportingIEsExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EventBasedReportingIEsExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EventBasedReportingIEsExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EventBasedReportingIEs-ExtIEsEntry", v)
}

// DecodeEventBasedReportingIEsExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEventBasedReportingIEsExtIEsEntry(data []byte) (EventBasedReportingIEsExtIEsEntry, int, error) {
	var v EventBasedReportingIEsExtIEsEntry
	err := ngapcodec.DecodeParameter("EventBasedReportingIEs-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EventBasedReportingIEsExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
