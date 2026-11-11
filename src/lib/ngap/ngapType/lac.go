// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LAC represents ASN.1 LAC.
type LAC struct {
	Value []byte `aper:"sizeLB:2,sizeUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LAC) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("LAC", v) }

// DecodeLAC decodes a standalone APER value and reports the consumed octets.
func DecodeLAC(data []byte) (LAC, int, error) {
	var v LAC
	err := ngapcodec.DecodeParameter("LAC", data, &v)
	if ngapcodec.IsFatal(err) {
		return LAC{}, 0, err
	}
	return v, len(data), err
}
