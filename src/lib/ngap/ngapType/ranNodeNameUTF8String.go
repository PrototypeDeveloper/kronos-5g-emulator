// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANNodeNameUTF8String represents ASN.1 RANNodeNameUTF8String.
type RANNodeNameUTF8String struct {
	Value string `aper:"sizeExt,sizeLB:1,sizeUB:150"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANNodeNameUTF8String) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RANNodeNameUTF8String", v)
}

// DecodeRANNodeNameUTF8String decodes a standalone APER value and reports the consumed octets.
func DecodeRANNodeNameUTF8String(data []byte) (RANNodeNameUTF8String, int, error) {
	var v RANNodeNameUTF8String
	err := ngapcodec.DecodeParameter("RANNodeNameUTF8String", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANNodeNameUTF8String{}, 0, err
	}
	return v, len(data), err
}
