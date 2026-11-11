// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// M5ReportAmountMDT represents ASN.1 M5ReportAmountMDT.
const (
	M5ReportAmountMDTPresentR1       Enumerated = 0
	M5ReportAmountMDTPresentR2       Enumerated = 1
	M5ReportAmountMDTPresentR4       Enumerated = 2
	M5ReportAmountMDTPresentR8       Enumerated = 3
	M5ReportAmountMDTPresentR16      Enumerated = 4
	M5ReportAmountMDTPresentR32      Enumerated = 5
	M5ReportAmountMDTPresentR64      Enumerated = 6
	M5ReportAmountMDTPresentInfinity Enumerated = 7
)

type M5ReportAmountMDT struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:7"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M5ReportAmountMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M5ReportAmountMDT", v)
}

// DecodeM5ReportAmountMDT decodes a standalone APER value and reports the consumed octets.
func DecodeM5ReportAmountMDT(data []byte) (M5ReportAmountMDT, int, error) {
	var v M5ReportAmountMDT
	err := ngapcodec.DecodeParameter("M5ReportAmountMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return M5ReportAmountMDT{}, 0, err
	}
	return v, len(data), err
}
