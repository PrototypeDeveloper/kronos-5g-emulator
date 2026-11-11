// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AllowedTACs represents ASN.1 AllowedTACs.
type AllowedTACs struct {
	List []TAC `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AllowedTACs) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("AllowedTACs", v) }

// DecodeAllowedTACs decodes a standalone APER value and reports the consumed octets.
func DecodeAllowedTACs(data []byte) (AllowedTACs, int, error) {
	var v AllowedTACs
	err := ngapcodec.DecodeParameter("AllowedTACs", data, &v)
	if ngapcodec.IsFatal(err) {
		return AllowedTACs{}, 0, err
	}
	return v, len(data), err
}
