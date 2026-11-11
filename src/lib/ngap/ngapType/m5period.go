// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// M5period represents ASN.1 M5period.
const (
	M5periodPresentMs1024  Enumerated = 0
	M5periodPresentMs2048  Enumerated = 1
	M5periodPresentMs5120  Enumerated = 2
	M5periodPresentMs10240 Enumerated = 3
	M5periodPresentMin1    Enumerated = 4
)

type M5period struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:4"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M5period) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("M5period", v) }

// DecodeM5period decodes a standalone APER value and reports the consumed octets.
func DecodeM5period(data []byte) (M5period, int, error) {
	var v M5period
	err := ngapcodec.DecodeParameter("M5period", data, &v)
	if ngapcodec.IsFatal(err) {
		return M5period{}, 0, err
	}
	return v, len(data), err
}
