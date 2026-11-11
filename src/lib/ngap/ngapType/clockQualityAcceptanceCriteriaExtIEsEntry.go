// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ClockQualityAcceptanceCriteriaExtIEsEntryPresent uint32

const (
	ClockQualityAcceptanceCriteriaExtIEsEntryPresentNothing ClockQualityAcceptanceCriteriaExtIEsEntryPresent = iota
	ClockQualityAcceptanceCriteriaExtIEsEntryPresentUnknown
)

type ClockQualityAcceptanceCriteriaExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      ClockQualityAcceptanceCriteriaExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ClockQualityAcceptanceCriteriaExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ClockQualityAcceptanceCriteria-ExtIEsEntry", v)
}

// DecodeClockQualityAcceptanceCriteriaExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeClockQualityAcceptanceCriteriaExtIEsEntry(data []byte) (ClockQualityAcceptanceCriteriaExtIEsEntry, int, error) {
	var v ClockQualityAcceptanceCriteriaExtIEsEntry
	err := ngapcodec.DecodeParameter("ClockQualityAcceptanceCriteria-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ClockQualityAcceptanceCriteriaExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
