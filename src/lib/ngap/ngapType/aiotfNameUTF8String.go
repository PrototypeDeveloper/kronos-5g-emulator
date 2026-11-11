// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIOTFNameUTF8String represents ASN.1 AIOTFNameUTF8String.
type AIOTFNameUTF8String struct {
	Value string `aper:"sizeExt,sizeLB:1,sizeUB:150"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIOTFNameUTF8String) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIOTFNameUTF8String", v)
}

// DecodeAIOTFNameUTF8String decodes a standalone APER value and reports the consumed octets.
func DecodeAIOTFNameUTF8String(data []byte) (AIOTFNameUTF8String, int, error) {
	var v AIOTFNameUTF8String
	err := ngapcodec.DecodeParameter("AIOTFNameUTF8String", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIOTFNameUTF8String{}, 0, err
	}
	return v, len(data), err
}
