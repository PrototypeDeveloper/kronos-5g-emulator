// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MDTModeEutra represents ASN.1 MDTModeEutra.
type MDTModeEutra struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTModeEutra) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("MDTModeEutra", v) }

// DecodeMDTModeEutra decodes a standalone APER value and reports the consumed octets.
func DecodeMDTModeEutra(data []byte) (MDTModeEutra, int, error) {
	var v MDTModeEutra
	err := ngapcodec.DecodeParameter("MDTModeEutra", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTModeEutra{}, 0, err
	}
	return v, len(data), err
}
