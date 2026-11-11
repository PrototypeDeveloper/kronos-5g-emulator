// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TNAPID represents ASN.1 TNAP-ID.
type TNAPID struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TNAPID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TNAP-ID", v) }

// DecodeTNAPID decodes a standalone APER value and reports the consumed octets.
func DecodeTNAPID(data []byte) (TNAPID, int, error) {
	var v TNAPID
	err := ngapcodec.DecodeParameter("TNAP-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return TNAPID{}, 0, err
	}
	return v, len(data), err
}
