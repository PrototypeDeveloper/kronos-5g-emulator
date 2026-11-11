// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RGLevelWirelineAccessCharacteristics represents ASN.1 RGLevelWirelineAccessCharacteristics.
type RGLevelWirelineAccessCharacteristics struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RGLevelWirelineAccessCharacteristics) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RGLevelWirelineAccessCharacteristics", v)
}

// DecodeRGLevelWirelineAccessCharacteristics decodes a standalone APER value and reports the consumed octets.
func DecodeRGLevelWirelineAccessCharacteristics(data []byte) (RGLevelWirelineAccessCharacteristics, int, error) {
	var v RGLevelWirelineAccessCharacteristics
	err := ngapcodec.DecodeParameter("RGLevelWirelineAccessCharacteristics", data, &v)
	if ngapcodec.IsFatal(err) {
		return RGLevelWirelineAccessCharacteristics{}, 0, err
	}
	return v, len(data), err
}
