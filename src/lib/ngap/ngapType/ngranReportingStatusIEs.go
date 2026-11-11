// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGRANReportingStatusIEs represents ASN.1 NGRAN-ReportingStatusIEs.
type NGRANReportingStatusIEs struct {
	NGRANCellReportList NGRANCellReportList
	HasIEExtensions     bool
	IEExtensions        []NGRANReportingStatusIEsExtIEsEntry
	UnknownExtensions   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANReportingStatusIEs) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGRAN-ReportingStatusIEs", v)
}

// DecodeNGRANReportingStatusIEs decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANReportingStatusIEs(data []byte) (NGRANReportingStatusIEs, int, error) {
	var v NGRANReportingStatusIEs
	err := ngapcodec.DecodeParameter("NGRAN-ReportingStatusIEs", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANReportingStatusIEs{}, 0, err
	}
	return v, len(data), err
}
