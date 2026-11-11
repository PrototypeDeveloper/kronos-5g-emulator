// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NSAGID represents ASN.1 NSAG-ID.
type NSAGID struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:255"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NSAGID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("NSAG-ID", v) }

// DecodeNSAGID decodes a standalone APER value and reports the consumed octets.
func DecodeNSAGID(data []byte) (NSAGID, int, error) {
	var v NSAGID
	err := ngapcodec.DecodeParameter("NSAG-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return NSAGID{}, 0, err
	}
	return v, len(data), err
}
