// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// InterSystemHOReport represents ASN.1 InterSystemHOReport.
type InterSystemHOReport struct {
	HandoverReportType InterSystemHandoverReportType
	HasIEExtensions    bool
	IEExtensions       []InterSystemHOReportExtIEsEntry
	UnknownExtensions  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InterSystemHOReport) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InterSystemHOReport", v)
}

// DecodeInterSystemHOReport decodes a standalone APER value and reports the consumed octets.
func DecodeInterSystemHOReport(data []byte) (InterSystemHOReport, int, error) {
	var v InterSystemHOReport
	err := ngapcodec.DecodeParameter("InterSystemHOReport", data, &v)
	if ngapcodec.IsFatal(err) {
		return InterSystemHOReport{}, 0, err
	}
	return v, len(data), err
}
