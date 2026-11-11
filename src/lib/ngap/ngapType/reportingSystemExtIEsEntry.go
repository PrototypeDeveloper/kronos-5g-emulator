// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ReportingSystemExtIEsEntryPresent uint32

const (
	ReportingSystemExtIEsEntryPresentNothing ReportingSystemExtIEsEntryPresent = iota
	ReportingSystemExtIEsEntryPresentUnknown
)

type ReportingSystemExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      ReportingSystemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ReportingSystemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ReportingSystem-ExtIEsEntry", v)
}

// DecodeReportingSystemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeReportingSystemExtIEsEntry(data []byte) (ReportingSystemExtIEsEntry, int, error) {
	var v ReportingSystemExtIEsEntry
	err := ngapcodec.DecodeParameter("ReportingSystem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ReportingSystemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
