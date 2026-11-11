// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAIListforQMC represents ASN.1 TAIListforQMC.
type TAIListforQMC struct {
	List []TAI `aper:"sizeLB:1,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIListforQMC) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TAIListforQMC", v) }

// DecodeTAIListforQMC decodes a standalone APER value and reports the consumed octets.
func DecodeTAIListforQMC(data []byte) (TAIListforQMC, int, error) {
	var v TAIListforQMC
	err := ngapcodec.DecodeParameter("TAIListforQMC", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIListforQMC{}, 0, err
	}
	return v, len(data), err
}
