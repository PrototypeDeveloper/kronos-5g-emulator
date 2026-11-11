// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIOTFNameVisibleString represents ASN.1 AIOTFNameVisibleString.
type AIOTFNameVisibleString struct {
	Value string `aper:"sizeExt,sizeLB:1,sizeUB:150"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIOTFNameVisibleString) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIOTFNameVisibleString", v)
}

// DecodeAIOTFNameVisibleString decodes a standalone APER value and reports the consumed octets.
func DecodeAIOTFNameVisibleString(data []byte) (AIOTFNameVisibleString, int, error) {
	var v AIOTFNameVisibleString
	err := ngapcodec.DecodeParameter("AIOTFNameVisibleString", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIOTFNameVisibleString{}, 0, err
	}
	return v, len(data), err
}
