// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MRBID represents ASN.1 MRB-ID.
type MRBID struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:512"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MRBID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("MRB-ID", v) }

// DecodeMRBID decodes a standalone APER value and reports the consumed octets.
func DecodeMRBID(data []byte) (MRBID, int, error) {
	var v MRBID
	err := ngapcodec.DecodeParameter("MRB-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return MRBID{}, 0, err
	}
	return v, len(data), err
}
