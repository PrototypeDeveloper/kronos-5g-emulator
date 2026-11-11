// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TWAPID represents ASN.1 TWAP-ID.
type TWAPID struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TWAPID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TWAP-ID", v) }

// DecodeTWAPID decodes a standalone APER value and reports the consumed octets.
func DecodeTWAPID(data []byte) (TWAPID, int, error) {
	var v TWAPID
	err := ngapcodec.DecodeParameter("TWAP-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return TWAPID{}, 0, err
	}
	return v, len(data), err
}
