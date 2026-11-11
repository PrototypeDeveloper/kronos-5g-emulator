// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExtendedConnectedTime represents ASN.1 Extended-ConnectedTime.
type ExtendedConnectedTime struct {
	Value int64 `aper:"valueLB:0,valueUB:255"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExtendedConnectedTime) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Extended-ConnectedTime", v)
}

// DecodeExtendedConnectedTime decodes a standalone APER value and reports the consumed octets.
func DecodeExtendedConnectedTime(data []byte) (ExtendedConnectedTime, int, error) {
	var v ExtendedConnectedTime
	err := ngapcodec.DecodeParameter("Extended-ConnectedTime", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExtendedConnectedTime{}, 0, err
	}
	return v, len(data), err
}
