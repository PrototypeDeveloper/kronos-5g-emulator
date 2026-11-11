// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExtendedReportIntervalMDT represents ASN.1 ExtendedReportIntervalMDT.
const (
	ExtendedReportIntervalMDTPresentMs20480 Enumerated = 0
	ExtendedReportIntervalMDTPresentMs40960 Enumerated = 1
)

type ExtendedReportIntervalMDT struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExtendedReportIntervalMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExtendedReportIntervalMDT", v)
}

// DecodeExtendedReportIntervalMDT decodes a standalone APER value and reports the consumed octets.
func DecodeExtendedReportIntervalMDT(data []byte) (ExtendedReportIntervalMDT, int, error) {
	var v ExtendedReportIntervalMDT
	err := ngapcodec.DecodeParameter("ExtendedReportIntervalMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExtendedReportIntervalMDT{}, 0, err
	}
	return v, len(data), err
}
