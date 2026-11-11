// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MaxNrofRSIndexesToReport represents ASN.1 MaxNrofRS-IndexesToReport.
type MaxNrofRSIndexesToReport struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MaxNrofRSIndexesToReport) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MaxNrofRS-IndexesToReport", v)
}

// DecodeMaxNrofRSIndexesToReport decodes a standalone APER value and reports the consumed octets.
func DecodeMaxNrofRSIndexesToReport(data []byte) (MaxNrofRSIndexesToReport, int, error) {
	var v MaxNrofRSIndexesToReport
	err := ngapcodec.DecodeParameter("MaxNrofRS-IndexesToReport", data, &v)
	if ngapcodec.IsFatal(err) {
		return MaxNrofRSIndexesToReport{}, 0, err
	}
	return v, len(data), err
}
