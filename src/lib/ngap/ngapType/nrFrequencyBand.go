// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRFrequencyBand represents ASN.1 NRFrequencyBand.
type NRFrequencyBand struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:1024"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRFrequencyBand) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NRFrequencyBand", v)
}

// DecodeNRFrequencyBand decodes a standalone APER value and reports the consumed octets.
func DecodeNRFrequencyBand(data []byte) (NRFrequencyBand, int, error) {
	var v NRFrequencyBand
	err := ngapcodec.DecodeParameter("NRFrequencyBand", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRFrequencyBand{}, 0, err
	}
	return v, len(data), err
}
