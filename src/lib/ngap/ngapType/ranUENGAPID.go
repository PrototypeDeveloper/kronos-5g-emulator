// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANUENGAPID represents ASN.1 RAN-UE-NGAP-ID.
type RANUENGAPID struct {
	Value int64 `aper:"valueLB:0,valueUB:4294967295"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANUENGAPID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("RAN-UE-NGAP-ID", v) }

// DecodeRANUENGAPID decodes a standalone APER value and reports the consumed octets.
func DecodeRANUENGAPID(data []byte) (RANUENGAPID, int, error) {
	var v RANUENGAPID
	err := ngapcodec.DecodeParameter("RAN-UE-NGAP-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANUENGAPID{}, 0, err
	}
	return v, len(data), err
}
