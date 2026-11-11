// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAC represents ASN.1 TAC.
type TAC struct {
	Value []byte `aper:"sizeLB:3,sizeUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAC) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TAC", v) }

// DecodeTAC decodes a standalone APER value and reports the consumed octets.
func DecodeTAC(data []byte) (TAC, int, error) {
	var v TAC
	err := ngapcodec.DecodeParameter("TAC", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAC{}, 0, err
	}
	return v, len(data), err
}
