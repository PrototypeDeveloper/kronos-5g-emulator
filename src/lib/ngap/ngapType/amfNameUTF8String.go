// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AMFNameUTF8String represents ASN.1 AMFNameUTF8String.
type AMFNameUTF8String struct {
	Value string `aper:"sizeExt,sizeLB:1,sizeUB:150"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFNameUTF8String) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AMFNameUTF8String", v)
}

// DecodeAMFNameUTF8String decodes a standalone APER value and reports the consumed octets.
func DecodeAMFNameUTF8String(data []byte) (AMFNameUTF8String, int, error) {
	var v AMFNameUTF8String
	err := ngapcodec.DecodeParameter("AMFNameUTF8String", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFNameUTF8String{}, 0, err
	}
	return v, len(data), err
}
