// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ClockQualityAcceptanceIndicationExtIEsEntryPresent uint32

const (
	ClockQualityAcceptanceIndicationExtIEsEntryPresentNothing ClockQualityAcceptanceIndicationExtIEsEntryPresent = iota
	ClockQualityAcceptanceIndicationExtIEsEntryPresentUnknown
)

type ClockQualityAcceptanceIndicationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      ClockQualityAcceptanceIndicationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ClockQualityAcceptanceIndicationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ClockQualityAcceptanceIndication-ExtIEsEntry", v)
}

// DecodeClockQualityAcceptanceIndicationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeClockQualityAcceptanceIndicationExtIEsEntry(data []byte) (ClockQualityAcceptanceIndicationExtIEsEntry, int, error) {
	var v ClockQualityAcceptanceIndicationExtIEsEntry
	err := ngapcodec.DecodeParameter("ClockQualityAcceptanceIndication-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ClockQualityAcceptanceIndicationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
