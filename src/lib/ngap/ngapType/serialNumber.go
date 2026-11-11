// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SerialNumber represents ASN.1 SerialNumber.
type SerialNumber struct {
	Value BitString `aper:"sizeLB:16,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SerialNumber) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("SerialNumber", v) }

// DecodeSerialNumber decodes a standalone APER value and reports the consumed octets.
func DecodeSerialNumber(data []byte) (SerialNumber, int, error) {
	var v SerialNumber
	err := ngapcodec.DecodeParameter("SerialNumber", data, &v)
	if ngapcodec.IsFatal(err) {
		return SerialNumber{}, 0, err
	}
	return v, len(data), err
}
