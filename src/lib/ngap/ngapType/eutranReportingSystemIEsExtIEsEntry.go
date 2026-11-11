// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EUTRANReportingSystemIEsExtIEsEntryPresent uint32

const (
	EUTRANReportingSystemIEsExtIEsEntryPresentNothing EUTRANReportingSystemIEsExtIEsEntryPresent = iota
	EUTRANReportingSystemIEsExtIEsEntryPresentUnknown
)

type EUTRANReportingSystemIEsExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EUTRANReportingSystemIEsExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRANReportingSystemIEsExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRAN-ReportingSystemIEs-ExtIEsEntry", v)
}

// DecodeEUTRANReportingSystemIEsExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRANReportingSystemIEsExtIEsEntry(data []byte) (EUTRANReportingSystemIEsExtIEsEntry, int, error) {
	var v EUTRANReportingSystemIEsExtIEsEntry
	err := ngapcodec.DecodeParameter("EUTRAN-ReportingSystemIEs-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRANReportingSystemIEsExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
