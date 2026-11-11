// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// Altitude represents ASN.1 Altitude.
type Altitude struct {
	Value int64 `aper:"valueExt,valueLB:-420,valueUB:10000"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v Altitude) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("Altitude", v) }

// DecodeAltitude decodes a standalone APER value and reports the consumed octets.
func DecodeAltitude(data []byte) (Altitude, int, error) {
	var v Altitude
	err := ngapcodec.DecodeParameter("Altitude", data, &v)
	if ngapcodec.IsFatal(err) {
		return Altitude{}, 0, err
	}
	return v, len(data), err
}
