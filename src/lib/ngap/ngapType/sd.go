// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SD represents ASN.1 SD.
type SD struct {
	Value []byte `aper:"sizeLB:3,sizeUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SD) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("SD", v) }

// DecodeSD decodes a standalone APER value and reports the consumed octets.
func DecodeSD(data []byte) (SD, int, error) {
	var v SD
	err := ngapcodec.DecodeParameter("SD", data, &v)
	if ngapcodec.IsFatal(err) {
		return SD{}, 0, err
	}
	return v, len(data), err
}
