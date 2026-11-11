// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntersystemSONInformationReport represents ASN.1 IntersystemSONInformationReport.
type IntersystemSONInformationReportPresent uint32

const (
	IntersystemSONInformationReportPresentNothing IntersystemSONInformationReportPresent = iota
	IntersystemSONInformationReportPresentHOReportInformation
	IntersystemSONInformationReportPresentFailureIndicationInformation
	IntersystemSONInformationReportPresentChoiceExtensions
)

type IntersystemSONInformationReport struct {
	Present                      IntersystemSONInformationReportPresent
	HOReportInformation          InterSystemHOReport
	FailureIndicationInformation InterSystemFailureIndication
	ChoiceExtensions             IntersystemSONInformationReportExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemSONInformationReport) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemSONInformationReport", v)
}

// DecodeIntersystemSONInformationReport decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemSONInformationReport(data []byte) (IntersystemSONInformationReport, int, error) {
	var v IntersystemSONInformationReport
	err := ngapcodec.DecodeParameter("IntersystemSONInformationReport", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemSONInformationReport{}, 0, err
	}
	return v, len(data), err
}
