// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ReportIntervalMDT represents ASN.1 ReportIntervalMDT.
const (
	ReportIntervalMDTPresentMs120   Enumerated = 0
	ReportIntervalMDTPresentMs240   Enumerated = 1
	ReportIntervalMDTPresentMs480   Enumerated = 2
	ReportIntervalMDTPresentMs640   Enumerated = 3
	ReportIntervalMDTPresentMs1024  Enumerated = 4
	ReportIntervalMDTPresentMs2048  Enumerated = 5
	ReportIntervalMDTPresentMs5120  Enumerated = 6
	ReportIntervalMDTPresentMs10240 Enumerated = 7
	ReportIntervalMDTPresentMin1    Enumerated = 8
	ReportIntervalMDTPresentMin6    Enumerated = 9
	ReportIntervalMDTPresentMin12   Enumerated = 10
	ReportIntervalMDTPresentMin30   Enumerated = 11
	ReportIntervalMDTPresentMin60   Enumerated = 12
)

type ReportIntervalMDT struct {
	Value Enumerated `aper:"valueLB:0,valueUB:12"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ReportIntervalMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ReportIntervalMDT", v)
}

// DecodeReportIntervalMDT decodes a standalone APER value and reports the consumed octets.
func DecodeReportIntervalMDT(data []byte) (ReportIntervalMDT, int, error) {
	var v ReportIntervalMDT
	err := ngapcodec.DecodeParameter("ReportIntervalMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return ReportIntervalMDT{}, 0, err
	}
	return v, len(data), err
}
