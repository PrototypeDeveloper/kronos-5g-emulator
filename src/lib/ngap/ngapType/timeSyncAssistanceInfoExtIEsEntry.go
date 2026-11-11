// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TimeSyncAssistanceInfoExtIEsEntryPresent uint32

const (
	TimeSyncAssistanceInfoExtIEsEntryPresentNothing TimeSyncAssistanceInfoExtIEsEntryPresent = iota
	TimeSyncAssistanceInfoExtIEsEntryPresentClockQualityReportingControlInfo
	TimeSyncAssistanceInfoExtIEsEntryPresentUnknown
)

type TimeSyncAssistanceInfoExtIEsEntry struct {
	ID                               ProtocolExtensionID
	Criticality                      Criticality
	Present                          TimeSyncAssistanceInfoExtIEsEntryPresent
	ClockQualityReportingControlInfo ClockQualityReportingControlInfo `ngap:"id:390,criticality:ignore,presence:optional"`
	UnknownValue                     []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TimeSyncAssistanceInfoExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TimeSyncAssistanceInfo-ExtIEsEntry", v)
}

// DecodeTimeSyncAssistanceInfoExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTimeSyncAssistanceInfoExtIEsEntry(data []byte) (TimeSyncAssistanceInfoExtIEsEntry, int, error) {
	var v TimeSyncAssistanceInfoExtIEsEntry
	err := ngapcodec.DecodeParameter("TimeSyncAssistanceInfo-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TimeSyncAssistanceInfoExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
