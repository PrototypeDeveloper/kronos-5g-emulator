// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SONInformationReport represents ASN.1 SONInformationReport.
type SONInformationReportPresent uint32

const (
	SONInformationReportPresentNothing SONInformationReportPresent = iota
	SONInformationReportPresentFailureIndicationInformation
	SONInformationReportPresentHOReportInformation
	SONInformationReportPresentChoiceExtensions
)

type SONInformationReport struct {
	Present                      SONInformationReportPresent
	FailureIndicationInformation FailureIndication
	HOReportInformation          HOReport
	ChoiceExtensions             SONInformationReportExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SONInformationReport) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SONInformationReport", v)
}

// DecodeSONInformationReport decodes a standalone APER value and reports the consumed octets.
func DecodeSONInformationReport(data []byte) (SONInformationReport, int, error) {
	var v SONInformationReport
	err := ngapcodec.DecodeParameter("SONInformationReport", data, &v)
	if ngapcodec.IsFatal(err) {
		return SONInformationReport{}, 0, err
	}
	return v, len(data), err
}
