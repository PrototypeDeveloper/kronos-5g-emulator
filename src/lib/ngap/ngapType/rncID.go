// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RNCID represents ASN.1 RNC-ID.
type RNCID struct {
	Value int64 `aper:"valueLB:0,valueUB:4095"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RNCID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("RNC-ID", v) }

// DecodeRNCID decodes a standalone APER value and reports the consumed octets.
func DecodeRNCID(data []byte) (RNCID, int, error) {
	var v RNCID
	err := ngapcodec.DecodeParameter("RNC-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return RNCID{}, 0, err
	}
	return v, len(data), err
}
