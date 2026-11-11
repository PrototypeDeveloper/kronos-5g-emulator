// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANNodeNameVisibleString represents ASN.1 RANNodeNameVisibleString.
type RANNodeNameVisibleString struct {
	Value string `aper:"sizeExt,sizeLB:1,sizeUB:150"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANNodeNameVisibleString) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RANNodeNameVisibleString", v)
}

// DecodeRANNodeNameVisibleString decodes a standalone APER value and reports the consumed octets.
func DecodeRANNodeNameVisibleString(data []byte) (RANNodeNameVisibleString, int, error) {
	var v RANNodeNameVisibleString
	err := ngapcodec.DecodeParameter("RANNodeNameVisibleString", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANNodeNameVisibleString{}, 0, err
	}
	return v, len(data), err
}
