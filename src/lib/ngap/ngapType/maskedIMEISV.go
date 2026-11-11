// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MaskedIMEISV represents ASN.1 MaskedIMEISV.
type MaskedIMEISV struct {
	Value BitString `aper:"sizeLB:64,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MaskedIMEISV) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("MaskedIMEISV", v) }

// DecodeMaskedIMEISV decodes a standalone APER value and reports the consumed octets.
func DecodeMaskedIMEISV(data []byte) (MaskedIMEISV, int, error) {
	var v MaskedIMEISV
	err := ngapcodec.DecodeParameter("MaskedIMEISV", data, &v)
	if ngapcodec.IsFatal(err) {
		return MaskedIMEISV{}, 0, err
	}
	return v, len(data), err
}
