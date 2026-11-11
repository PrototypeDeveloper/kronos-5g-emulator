// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ReportTypeExtIEsEntryPresent uint32

const (
	ReportTypeExtIEsEntryPresentNothing ReportTypeExtIEsEntryPresent = iota
	ReportTypeExtIEsEntryPresentUnknown
)

type ReportTypeExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      ReportTypeExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ReportTypeExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ReportType-ExtIEsEntry", v)
}

// DecodeReportTypeExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeReportTypeExtIEsEntry(data []byte) (ReportTypeExtIEsEntry, int, error) {
	var v ReportTypeExtIEsEntry
	err := ngapcodec.DecodeParameter("ReportType-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ReportTypeExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
