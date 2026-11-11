// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// M4period represents ASN.1 M4period.
const (
	M4periodPresentMs1024  Enumerated = 0
	M4periodPresentMs2048  Enumerated = 1
	M4periodPresentMs5120  Enumerated = 2
	M4periodPresentMs10240 Enumerated = 3
	M4periodPresentMin1    Enumerated = 4
)

type M4period struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:4"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M4period) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("M4period", v) }

// DecodeM4period decodes a standalone APER value and reports the consumed octets.
func DecodeM4period(data []byte) (M4period, int, error) {
	var v M4period
	err := ngapcodec.DecodeParameter("M4period", data, &v)
	if ngapcodec.IsFatal(err) {
		return M4period{}, 0, err
	}
	return v, len(data), err
}
