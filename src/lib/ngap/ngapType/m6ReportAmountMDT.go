// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// M6ReportAmountMDT represents ASN.1 M6ReportAmountMDT.
const (
	M6ReportAmountMDTPresentR1       Enumerated = 0
	M6ReportAmountMDTPresentR2       Enumerated = 1
	M6ReportAmountMDTPresentR4       Enumerated = 2
	M6ReportAmountMDTPresentR8       Enumerated = 3
	M6ReportAmountMDTPresentR16      Enumerated = 4
	M6ReportAmountMDTPresentR32      Enumerated = 5
	M6ReportAmountMDTPresentR64      Enumerated = 6
	M6ReportAmountMDTPresentInfinity Enumerated = 7
)

type M6ReportAmountMDT struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:7"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M6ReportAmountMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M6ReportAmountMDT", v)
}

// DecodeM6ReportAmountMDT decodes a standalone APER value and reports the consumed octets.
func DecodeM6ReportAmountMDT(data []byte) (M6ReportAmountMDT, int, error) {
	var v M6ReportAmountMDT
	err := ngapcodec.DecodeParameter("M6ReportAmountMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return M6ReportAmountMDT{}, 0, err
	}
	return v, len(data), err
}
