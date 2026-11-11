// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAListforMDT represents ASN.1 TAListforMDT.
type TAListforMDT struct {
	List []TAC `aper:"sizeLB:1,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAListforMDT) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TAListforMDT", v) }

// DecodeTAListforMDT decodes a standalone APER value and reports the consumed octets.
func DecodeTAListforMDT(data []byte) (TAListforMDT, int, error) {
	var v TAListforMDT
	err := ngapcodec.DecodeParameter("TAListforMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAListforMDT{}, 0, err
	}
	return v, len(data), err
}
