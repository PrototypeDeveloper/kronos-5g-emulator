// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// M4ReportAmountMDT represents ASN.1 M4ReportAmountMDT.
const (
	M4ReportAmountMDTPresentR1       Enumerated = 0
	M4ReportAmountMDTPresentR2       Enumerated = 1
	M4ReportAmountMDTPresentR4       Enumerated = 2
	M4ReportAmountMDTPresentR8       Enumerated = 3
	M4ReportAmountMDTPresentR16      Enumerated = 4
	M4ReportAmountMDTPresentR32      Enumerated = 5
	M4ReportAmountMDTPresentR64      Enumerated = 6
	M4ReportAmountMDTPresentInfinity Enumerated = 7
)

type M4ReportAmountMDT struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:7"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M4ReportAmountMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M4ReportAmountMDT", v)
}

// DecodeM4ReportAmountMDT decodes a standalone APER value and reports the consumed octets.
func DecodeM4ReportAmountMDT(data []byte) (M4ReportAmountMDT, int, error) {
	var v M4ReportAmountMDT
	err := ngapcodec.DecodeParameter("M4ReportAmountMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return M4ReportAmountMDT{}, 0, err
	}
	return v, len(data), err
}
