// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TMGI represents ASN.1 TMGI.
type TMGI struct {
	Value []byte `aper:"sizeLB:6,sizeUB:6"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TMGI) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TMGI", v) }

// DecodeTMGI decodes a standalone APER value and reports the consumed octets.
func DecodeTMGI(data []byte) (TMGI, int, error) {
	var v TMGI
	err := ngapcodec.DecodeParameter("TMGI", data, &v)
	if ngapcodec.IsFatal(err) {
		return TMGI{}, 0, err
	}
	return v, len(data), err
}
