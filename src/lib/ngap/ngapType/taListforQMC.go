// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAListforQMC represents ASN.1 TAListforQMC.
type TAListforQMC struct {
	List []TAC `aper:"sizeLB:1,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAListforQMC) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TAListforQMC", v) }

// DecodeTAListforQMC decodes a standalone APER value and reports the consumed octets.
func DecodeTAListforQMC(data []byte) (TAListforQMC, int, error) {
	var v TAListforQMC
	err := ngapcodec.DecodeParameter("TAListforQMC", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAListforQMC{}, 0, err
	}
	return v, len(data), err
}
