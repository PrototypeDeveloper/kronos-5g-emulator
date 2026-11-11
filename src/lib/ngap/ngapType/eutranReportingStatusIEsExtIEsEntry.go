// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EUTRANReportingStatusIEsExtIEsEntryPresent uint32

const (
	EUTRANReportingStatusIEsExtIEsEntryPresentNothing EUTRANReportingStatusIEsExtIEsEntryPresent = iota
	EUTRANReportingStatusIEsExtIEsEntryPresentUnknown
)

type EUTRANReportingStatusIEsExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EUTRANReportingStatusIEsExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRANReportingStatusIEsExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRAN-ReportingStatusIEs-ExtIEsEntry", v)
}

// DecodeEUTRANReportingStatusIEsExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRANReportingStatusIEsExtIEsEntry(data []byte) (EUTRANReportingStatusIEsExtIEsEntry, int, error) {
	var v EUTRANReportingStatusIEsExtIEsEntry
	err := ngapcodec.DecodeParameter("EUTRAN-ReportingStatusIEs-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRANReportingStatusIEsExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
