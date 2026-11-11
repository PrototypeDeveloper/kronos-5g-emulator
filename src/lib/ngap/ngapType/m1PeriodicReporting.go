// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// M1PeriodicReporting represents ASN.1 M1PeriodicReporting.
type M1PeriodicReporting struct {
	ReportInterval    ReportIntervalMDT
	ReportAmount      ReportAmountMDT
	HasIEExtensions   bool
	IEExtensions      []M1PeriodicReportingExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M1PeriodicReporting) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M1PeriodicReporting", v)
}

// DecodeM1PeriodicReporting decodes a standalone APER value and reports the consumed octets.
func DecodeM1PeriodicReporting(data []byte) (M1PeriodicReporting, int, error) {
	var v M1PeriodicReporting
	err := ngapcodec.DecodeParameter("M1PeriodicReporting", data, &v)
	if ngapcodec.IsFatal(err) {
		return M1PeriodicReporting{}, 0, err
	}
	return v, len(data), err
}
