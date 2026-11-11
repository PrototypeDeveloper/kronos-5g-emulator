// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// Periodicity represents ASN.1 Periodicity.
type Periodicity struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:640000"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v Periodicity) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("Periodicity", v) }

// DecodePeriodicity decodes a standalone APER value and reports the consumed octets.
func DecodePeriodicity(data []byte) (Periodicity, int, error) {
	var v Periodicity
	err := ngapcodec.DecodeParameter("Periodicity", data, &v)
	if ngapcodec.IsFatal(err) {
		return Periodicity{}, 0, err
	}
	return v, len(data), err
}
