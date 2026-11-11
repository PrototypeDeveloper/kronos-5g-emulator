// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// HOReport represents ASN.1 HOReport.
type HOReport struct {
	HandoverReportType        HOReportHandoverReportType `aper:"valueExt,valueLB:0,valueUB:2"`
	HandoverCause             Cause
	SourcecellCGI             NGRANCGI
	TargetcellCGI             NGRANCGI
	HasReestablishmentcellCGI bool
	ReestablishmentcellCGI    NGRANCGI
	HasSourcecellCRNTI        bool
	SourcecellCRNTI           BitString `aper:"sizeLB:16,sizeUB:16"`
	HasTargetcellinEUTRAN     bool
	TargetcellinEUTRAN        EUTRACGI
	HasMobilityInformation    bool
	MobilityInformation       MobilityInformation
	HasUERLFReportContainer   bool
	UERLFReportContainer      UERLFReportContainer
	HasIEExtensions           bool
	IEExtensions              []HOReportExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HOReport) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("HOReport", v) }

// DecodeHOReport decodes a standalone APER value and reports the consumed octets.
func DecodeHOReport(data []byte) (HOReport, int, error) {
	var v HOReport
	err := ngapcodec.DecodeParameter("HOReport", data, &v)
	if ngapcodec.IsFatal(err) {
		return HOReport{}, 0, err
	}
	return v, len(data), err
}
