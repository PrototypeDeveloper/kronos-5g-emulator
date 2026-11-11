// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AMFPointer represents ASN.1 AMFPointer.
type AMFPointer struct {
	Value BitString `aper:"sizeLB:6,sizeUB:6"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFPointer) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("AMFPointer", v) }

// DecodeAMFPointer decodes a standalone APER value and reports the consumed octets.
func DecodeAMFPointer(data []byte) (AMFPointer, int, error) {
	var v AMFPointer
	err := ngapcodec.DecodeParameter("AMFPointer", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFPointer{}, 0, err
	}
	return v, len(data), err
}
