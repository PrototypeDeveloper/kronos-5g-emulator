// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ReportType represents ASN.1 ReportType.
type ReportTypePresent uint32

const (
	ReportTypePresentNothing ReportTypePresent = iota
	ReportTypePresentEventBasedReporting
	ReportTypePresentPeriodicReporting
	ReportTypePresentChoiceExtensions
)

type ReportType struct {
	Present             ReportTypePresent
	EventBasedReporting EventBasedReportingIEs
	PeriodicReporting   PeriodicReportingIEs
	ChoiceExtensions    ReportTypeExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ReportType) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("ReportType", v) }

// DecodeReportType decodes a standalone APER value and reports the consumed octets.
func DecodeReportType(data []byte) (ReportType, int, error) {
	var v ReportType
	err := ngapcodec.DecodeParameter("ReportType", data, &v)
	if ngapcodec.IsFatal(err) {
		return ReportType{}, 0, err
	}
	return v, len(data), err
}
