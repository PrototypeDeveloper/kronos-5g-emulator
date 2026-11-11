// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ERABID represents ASN.1 E-RAB-ID.
type ERABID struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:15"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ERABID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("E-RAB-ID", v) }

// DecodeERABID decodes a standalone APER value and reports the consumed octets.
func DecodeERABID(data []byte) (ERABID, int, error) {
	var v ERABID
	err := ngapcodec.DecodeParameter("E-RAB-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return ERABID{}, 0, err
	}
	return v, len(data), err
}
