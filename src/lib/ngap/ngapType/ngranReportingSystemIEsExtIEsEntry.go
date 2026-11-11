// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NGRANReportingSystemIEsExtIEsEntryPresent uint32

const (
	NGRANReportingSystemIEsExtIEsEntryPresentNothing NGRANReportingSystemIEsExtIEsEntryPresent = iota
	NGRANReportingSystemIEsExtIEsEntryPresentUnknown
)

type NGRANReportingSystemIEsExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NGRANReportingSystemIEsExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANReportingSystemIEsExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGRAN-ReportingSystemIEs-ExtIEsEntry", v)
}

// DecodeNGRANReportingSystemIEsExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANReportingSystemIEsExtIEsEntry(data []byte) (NGRANReportingSystemIEsExtIEsEntry, int, error) {
	var v NGRANReportingSystemIEsExtIEsEntry
	err := ngapcodec.DecodeParameter("NGRAN-ReportingSystemIEs-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANReportingSystemIEsExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
