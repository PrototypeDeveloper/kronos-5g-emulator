// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SgNBUEX2APID represents ASN.1 SgNB-UE-X2AP-ID.
type SgNBUEX2APID struct {
	Value int64 `aper:"valueLB:0,valueUB:4294967295"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SgNBUEX2APID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SgNB-UE-X2AP-ID", v)
}

// DecodeSgNBUEX2APID decodes a standalone APER value and reports the consumed octets.
func DecodeSgNBUEX2APID(data []byte) (SgNBUEX2APID, int, error) {
	var v SgNBUEX2APID
	err := ngapcodec.DecodeParameter("SgNB-UE-X2AP-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return SgNBUEX2APID{}, 0, err
	}
	return v, len(data), err
}
