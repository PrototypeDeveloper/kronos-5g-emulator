// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExtendedRNCID represents ASN.1 ExtendedRNC-ID.
type ExtendedRNCID struct {
	Value int64 `aper:"valueLB:4096,valueUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExtendedRNCID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExtendedRNC-ID", v)
}

// DecodeExtendedRNCID decodes a standalone APER value and reports the consumed octets.
func DecodeExtendedRNCID(data []byte) (ExtendedRNCID, int, error) {
	var v ExtendedRNCID
	err := ngapcodec.DecodeParameter("ExtendedRNC-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExtendedRNCID{}, 0, err
	}
	return v, len(data), err
}
