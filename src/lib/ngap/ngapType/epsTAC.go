// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EPSTAC represents ASN.1 EPS-TAC.
type EPSTAC struct {
	Value []byte `aper:"sizeLB:2,sizeUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EPSTAC) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("EPS-TAC", v) }

// DecodeEPSTAC decodes a standalone APER value and reports the consumed octets.
func DecodeEPSTAC(data []byte) (EPSTAC, int, error) {
	var v EPSTAC
	err := ngapcodec.DecodeParameter("EPS-TAC", data, &v)
	if ngapcodec.IsFatal(err) {
		return EPSTAC{}, 0, err
	}
	return v, len(data), err
}
