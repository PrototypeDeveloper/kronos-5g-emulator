// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// M7ReportAmountMDT represents ASN.1 M7ReportAmountMDT.
const (
	M7ReportAmountMDTPresentR1       Enumerated = 0
	M7ReportAmountMDTPresentR2       Enumerated = 1
	M7ReportAmountMDTPresentR4       Enumerated = 2
	M7ReportAmountMDTPresentR8       Enumerated = 3
	M7ReportAmountMDTPresentR16      Enumerated = 4
	M7ReportAmountMDTPresentR32      Enumerated = 5
	M7ReportAmountMDTPresentR64      Enumerated = 6
	M7ReportAmountMDTPresentInfinity Enumerated = 7
)

type M7ReportAmountMDT struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:7"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M7ReportAmountMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M7ReportAmountMDT", v)
}

// DecodeM7ReportAmountMDT decodes a standalone APER value and reports the consumed octets.
func DecodeM7ReportAmountMDT(data []byte) (M7ReportAmountMDT, int, error) {
	var v M7ReportAmountMDT
	err := ngapcodec.DecodeParameter("M7ReportAmountMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return M7ReportAmountMDT{}, 0, err
	}
	return v, len(data), err
}
