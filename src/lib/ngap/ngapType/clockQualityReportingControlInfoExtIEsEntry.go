// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ClockQualityReportingControlInfoExtIEsEntryPresent uint32

const (
	ClockQualityReportingControlInfoExtIEsEntryPresentNothing ClockQualityReportingControlInfoExtIEsEntryPresent = iota
	ClockQualityReportingControlInfoExtIEsEntryPresentUnknown
)

type ClockQualityReportingControlInfoExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      ClockQualityReportingControlInfoExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ClockQualityReportingControlInfoExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ClockQualityReportingControlInfo-ExtIEsEntry", v)
}

// DecodeClockQualityReportingControlInfoExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeClockQualityReportingControlInfoExtIEsEntry(data []byte) (ClockQualityReportingControlInfoExtIEsEntry, int, error) {
	var v ClockQualityReportingControlInfoExtIEsEntry
	err := ngapcodec.DecodeParameter("ClockQualityReportingControlInfo-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ClockQualityReportingControlInfoExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
