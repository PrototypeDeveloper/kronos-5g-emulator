// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAIListforMDT represents ASN.1 TAIListforMDT.
type TAIListforMDT struct {
	List []TAI `aper:"sizeLB:1,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIListforMDT) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TAIListforMDT", v) }

// DecodeTAIListforMDT decodes a standalone APER value and reports the consumed octets.
func DecodeTAIListforMDT(data []byte) (TAIListforMDT, int, error) {
	var v TAIListforMDT
	err := ngapcodec.DecodeParameter("TAIListforMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIListforMDT{}, 0, err
	}
	return v, len(data), err
}
