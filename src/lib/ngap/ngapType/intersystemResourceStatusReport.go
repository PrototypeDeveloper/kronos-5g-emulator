// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntersystemResourceStatusReport represents ASN.1 IntersystemResourceStatusReport.
type IntersystemResourceStatusReport struct {
	ReportingSystem   ResourceStatusReportingSystem
	HasIEExtensions   bool
	IEExtensions      []IntersystemResourceStatusReportExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemResourceStatusReport) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemResourceStatusReport", v)
}

// DecodeIntersystemResourceStatusReport decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemResourceStatusReport(data []byte) (IntersystemResourceStatusReport, int, error) {
	var v IntersystemResourceStatusReport
	err := ngapcodec.DecodeParameter("IntersystemResourceStatusReport", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemResourceStatusReport{}, 0, err
	}
	return v, len(data), err
}
