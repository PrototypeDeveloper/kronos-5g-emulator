// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGRANReportingSystemIEs represents ASN.1 NGRAN-ReportingSystemIEs.
type NGRANReportingSystemIEs struct {
	NGRANCellToReportList NGRANCellToReportList
	HasIEExtensions       bool
	IEExtensions          []NGRANReportingSystemIEsExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANReportingSystemIEs) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGRAN-ReportingSystemIEs", v)
}

// DecodeNGRANReportingSystemIEs decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANReportingSystemIEs(data []byte) (NGRANReportingSystemIEs, int, error) {
	var v NGRANReportingSystemIEs
	err := ngapcodec.DecodeParameter("NGRAN-ReportingSystemIEs", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANReportingSystemIEs{}, 0, err
	}
	return v, len(data), err
}
