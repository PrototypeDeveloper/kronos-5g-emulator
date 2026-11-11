// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PeriodicReportingIEsExtIEsEntryPresent uint32

const (
	PeriodicReportingIEsExtIEsEntryPresentNothing PeriodicReportingIEsExtIEsEntryPresent = iota
	PeriodicReportingIEsExtIEsEntryPresentUnknown
)

type PeriodicReportingIEsExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PeriodicReportingIEsExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PeriodicReportingIEsExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PeriodicReportingIEs-ExtIEsEntry", v)
}

// DecodePeriodicReportingIEsExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePeriodicReportingIEsExtIEsEntry(data []byte) (PeriodicReportingIEsExtIEsEntry, int, error) {
	var v PeriodicReportingIEsExtIEsEntry
	err := ngapcodec.DecodeParameter("PeriodicReportingIEs-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PeriodicReportingIEsExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
