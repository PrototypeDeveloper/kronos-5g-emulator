// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EUTRANReportingSystemIEs represents ASN.1 EUTRAN-ReportingSystemIEs.
type EUTRANReportingSystemIEs struct {
	EUTRANCellToReportList EUTRANCellToReportList
	HasIEExtensions        bool
	IEExtensions           []EUTRANReportingSystemIEsExtIEsEntry
	UnknownExtensions      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRANReportingSystemIEs) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRAN-ReportingSystemIEs", v)
}

// DecodeEUTRANReportingSystemIEs decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRANReportingSystemIEs(data []byte) (EUTRANReportingSystemIEs, int, error) {
	var v EUTRANReportingSystemIEs
	err := ngapcodec.DecodeParameter("EUTRAN-ReportingSystemIEs", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRANReportingSystemIEs{}, 0, err
	}
	return v, len(data), err
}
