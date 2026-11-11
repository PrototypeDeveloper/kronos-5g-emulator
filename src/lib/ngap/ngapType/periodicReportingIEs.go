// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PeriodicReportingIEs represents ASN.1 PeriodicReportingIEs.
type PeriodicReportingIEs struct {
	ReportingPeriodicity ReportingPeriodicity
	HasIEExtensions      bool
	IEExtensions         []PeriodicReportingIEsExtIEsEntry
	UnknownExtensions    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PeriodicReportingIEs) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PeriodicReportingIEs", v)
}

// DecodePeriodicReportingIEs decodes a standalone APER value and reports the consumed octets.
func DecodePeriodicReportingIEs(data []byte) (PeriodicReportingIEs, int, error) {
	var v PeriodicReportingIEs
	err := ngapcodec.DecodeParameter("PeriodicReportingIEs", data, &v)
	if ngapcodec.IsFatal(err) {
		return PeriodicReportingIEs{}, 0, err
	}
	return v, len(data), err
}
