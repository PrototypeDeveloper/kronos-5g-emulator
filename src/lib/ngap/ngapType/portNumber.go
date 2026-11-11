// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PortNumber represents ASN.1 PortNumber.
type PortNumber struct {
	Value []byte `aper:"sizeLB:2,sizeUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PortNumber) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("PortNumber", v) }

// DecodePortNumber decodes a standalone APER value and reports the consumed octets.
func DecodePortNumber(data []byte) (PortNumber, int, error) {
	var v PortNumber
	err := ngapcodec.DecodeParameter("PortNumber", data, &v)
	if ngapcodec.IsFatal(err) {
		return PortNumber{}, 0, err
	}
	return v, len(data), err
}
