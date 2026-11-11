// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ForbiddenTACs represents ASN.1 ForbiddenTACs.
type ForbiddenTACs struct {
	List []TAC `aper:"sizeLB:1,sizeUB:4096"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ForbiddenTACs) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("ForbiddenTACs", v) }

// DecodeForbiddenTACs decodes a standalone APER value and reports the consumed octets.
func DecodeForbiddenTACs(data []byte) (ForbiddenTACs, int, error) {
	var v ForbiddenTACs
	err := ngapcodec.DecodeParameter("ForbiddenTACs", data, &v)
	if ngapcodec.IsFatal(err) {
		return ForbiddenTACs{}, 0, err
	}
	return v, len(data), err
}
