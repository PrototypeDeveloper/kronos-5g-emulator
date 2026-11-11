// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AMFNameVisibleString represents ASN.1 AMFNameVisibleString.
type AMFNameVisibleString struct {
	Value string `aper:"sizeExt,sizeLB:1,sizeUB:150"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFNameVisibleString) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AMFNameVisibleString", v)
}

// DecodeAMFNameVisibleString decodes a standalone APER value and reports the consumed octets.
func DecodeAMFNameVisibleString(data []byte) (AMFNameVisibleString, int, error) {
	var v AMFNameVisibleString
	err := ngapcodec.DecodeParameter("AMFNameVisibleString", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFNameVisibleString{}, 0, err
	}
	return v, len(data), err
}
