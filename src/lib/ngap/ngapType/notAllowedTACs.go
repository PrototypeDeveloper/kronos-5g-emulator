// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NotAllowedTACs represents ASN.1 NotAllowedTACs.
type NotAllowedTACs struct {
	List []TAC `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NotAllowedTACs) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NotAllowedTACs", v)
}

// DecodeNotAllowedTACs decodes a standalone APER value and reports the consumed octets.
func DecodeNotAllowedTACs(data []byte) (NotAllowedTACs, int, error) {
	var v NotAllowedTACs
	err := ngapcodec.DecodeParameter("NotAllowedTACs", data, &v)
	if ngapcodec.IsFatal(err) {
		return NotAllowedTACs{}, 0, err
	}
	return v, len(data), err
}
