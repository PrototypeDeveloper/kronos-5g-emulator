// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MaximumIntegrityProtectedDataRate represents ASN.1 MaximumIntegrityProtectedDataRate.
const (
	MaximumIntegrityProtectedDataRatePresentBitrate64kbs  Enumerated = 0
	MaximumIntegrityProtectedDataRatePresentMaximumUERate Enumerated = 1
)

type MaximumIntegrityProtectedDataRate struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MaximumIntegrityProtectedDataRate) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MaximumIntegrityProtectedDataRate", v)
}

// DecodeMaximumIntegrityProtectedDataRate decodes a standalone APER value and reports the consumed octets.
func DecodeMaximumIntegrityProtectedDataRate(data []byte) (MaximumIntegrityProtectedDataRate, int, error) {
	var v MaximumIntegrityProtectedDataRate
	err := ngapcodec.DecodeParameter("MaximumIntegrityProtectedDataRate", data, &v)
	if ngapcodec.IsFatal(err) {
		return MaximumIntegrityProtectedDataRate{}, 0, err
	}
	return v, len(data), err
}
