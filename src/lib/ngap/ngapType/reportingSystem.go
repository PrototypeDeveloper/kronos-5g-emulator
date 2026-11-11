// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ReportingSystem represents ASN.1 ReportingSystem.
type ReportingSystemPresent uint32

const (
	ReportingSystemPresentNothing ReportingSystemPresent = iota
	ReportingSystemPresentEUTRAN
	ReportingSystemPresentNGRAN
	ReportingSystemPresentNoReporting
	ReportingSystemPresentChoiceExtensions
)

type ReportingSystem struct {
	Present          ReportingSystemPresent
	EUTRAN           EUTRANReportingSystemIEs
	NGRAN            NGRANReportingSystemIEs
	NoReporting      struct{}
	ChoiceExtensions ReportingSystemExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ReportingSystem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ReportingSystem", v)
}

// DecodeReportingSystem decodes a standalone APER value and reports the consumed octets.
func DecodeReportingSystem(data []byte) (ReportingSystem, int, error) {
	var v ReportingSystem
	err := ngapcodec.DecodeParameter("ReportingSystem", data, &v)
	if ngapcodec.IsFatal(err) {
		return ReportingSystem{}, 0, err
	}
	return v, len(data), err
}
