// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CAGID represents ASN.1 CAG-ID.
type CAGID struct {
	Value BitString `aper:"sizeLB:32,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CAGID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("CAG-ID", v) }

// DecodeCAGID decodes a standalone APER value and reports the consumed octets.
func DecodeCAGID(data []byte) (CAGID, int, error) {
	var v CAGID
	err := ngapcodec.DecodeParameter("CAG-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return CAGID{}, 0, err
	}
	return v, len(data), err
}
