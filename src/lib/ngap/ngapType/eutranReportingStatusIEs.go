// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EUTRANReportingStatusIEs represents ASN.1 EUTRAN-ReportingStatusIEs.
type EUTRANReportingStatusIEs struct {
	EUTRANCellReportList EUTRANCellReportList
	HasIEExtensions      bool
	IEExtensions         []EUTRANReportingStatusIEsExtIEsEntry
	UnknownExtensions    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRANReportingStatusIEs) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRAN-ReportingStatusIEs", v)
}

// DecodeEUTRANReportingStatusIEs decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRANReportingStatusIEs(data []byte) (EUTRANReportingStatusIEs, int, error) {
	var v EUTRANReportingStatusIEs
	err := ngapcodec.DecodeParameter("EUTRAN-ReportingStatusIEs", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRANReportingStatusIEs{}, 0, err
	}
	return v, len(data), err
}
