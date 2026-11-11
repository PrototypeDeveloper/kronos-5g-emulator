// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DRBID represents ASN.1 DRB-ID.
type DRBID struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("DRB-ID", v) }

// DecodeDRBID decodes a standalone APER value and reports the consumed octets.
func DecodeDRBID(data []byte) (DRBID, int, error) {
	var v DRBID
	err := ngapcodec.DecodeParameter("DRB-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBID{}, 0, err
	}
	return v, len(data), err
}
