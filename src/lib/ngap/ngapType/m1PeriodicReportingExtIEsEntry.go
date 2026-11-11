// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type M1PeriodicReportingExtIEsEntryPresent uint32

const (
	M1PeriodicReportingExtIEsEntryPresentNothing M1PeriodicReportingExtIEsEntryPresent = iota
	M1PeriodicReportingExtIEsEntryPresentExtendedReportIntervalMDT
	M1PeriodicReportingExtIEsEntryPresentUnknown
)

type M1PeriodicReportingExtIEsEntry struct {
	ID                        ProtocolExtensionID
	Criticality               Criticality
	Present                   M1PeriodicReportingExtIEsEntryPresent
	ExtendedReportIntervalMDT ExtendedReportIntervalMDT `ngap:"id:285,criticality:ignore,presence:optional"`
	UnknownValue              []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M1PeriodicReportingExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M1PeriodicReporting-ExtIEsEntry", v)
}

// DecodeM1PeriodicReportingExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeM1PeriodicReportingExtIEsEntry(data []byte) (M1PeriodicReportingExtIEsEntry, int, error) {
	var v M1PeriodicReportingExtIEsEntry
	err := ngapcodec.DecodeParameter("M1PeriodicReporting-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return M1PeriodicReportingExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
