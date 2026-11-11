// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAICancelledEUTRA represents ASN.1 TAICancelledEUTRA.
type TAICancelledEUTRA struct {
	List []TAICancelledEUTRAItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAICancelledEUTRA) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAICancelledEUTRA", v)
}

// DecodeTAICancelledEUTRA decodes a standalone APER value and reports the consumed octets.
func DecodeTAICancelledEUTRA(data []byte) (TAICancelledEUTRA, int, error) {
	var v TAICancelledEUTRA
	err := ngapcodec.DecodeParameter("TAICancelledEUTRA", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAICancelledEUTRA{}, 0, err
	}
	return v, len(data), err
}
