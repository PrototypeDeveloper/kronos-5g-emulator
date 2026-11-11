// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ReportAmountMDT represents ASN.1 ReportAmountMDT.
const (
	ReportAmountMDTPresentR1        Enumerated = 0
	ReportAmountMDTPresentR2        Enumerated = 1
	ReportAmountMDTPresentR4        Enumerated = 2
	ReportAmountMDTPresentR8        Enumerated = 3
	ReportAmountMDTPresentR16       Enumerated = 4
	ReportAmountMDTPresentR32       Enumerated = 5
	ReportAmountMDTPresentR64       Enumerated = 6
	ReportAmountMDTPresentRinfinity Enumerated = 7
)

type ReportAmountMDT struct {
	Value Enumerated `aper:"valueLB:0,valueUB:7"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ReportAmountMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ReportAmountMDT", v)
}

// DecodeReportAmountMDT decodes a standalone APER value and reports the consumed octets.
func DecodeReportAmountMDT(data []byte) (ReportAmountMDT, int, error) {
	var v ReportAmountMDT
	err := ngapcodec.DecodeParameter("ReportAmountMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return ReportAmountMDT{}, 0, err
	}
	return v, len(data), err
}
