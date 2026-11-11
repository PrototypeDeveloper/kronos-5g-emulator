// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NGRANReportingStatusIEsExtIEsEntryPresent uint32

const (
	NGRANReportingStatusIEsExtIEsEntryPresentNothing NGRANReportingStatusIEsExtIEsEntryPresent = iota
	NGRANReportingStatusIEsExtIEsEntryPresentUnknown
)

type NGRANReportingStatusIEsExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NGRANReportingStatusIEsExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANReportingStatusIEsExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGRAN-ReportingStatusIEs-ExtIEsEntry", v)
}

// DecodeNGRANReportingStatusIEsExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANReportingStatusIEsExtIEsEntry(data []byte) (NGRANReportingStatusIEsExtIEsEntry, int, error) {
	var v NGRANReportingStatusIEsExtIEsEntry
	err := ngapcodec.DecodeParameter("NGRAN-ReportingStatusIEs-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANReportingStatusIEsExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
