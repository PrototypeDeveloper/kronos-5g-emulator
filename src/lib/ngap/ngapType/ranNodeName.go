// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANNodeName represents ASN.1 RANNodeName.
type RANNodeName struct {
	Value string `aper:"sizeExt,sizeLB:1,sizeUB:150"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANNodeName) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("RANNodeName", v) }

// DecodeRANNodeName decodes a standalone APER value and reports the consumed octets.
func DecodeRANNodeName(data []byte) (RANNodeName, int, error) {
	var v RANNodeName
	err := ngapcodec.DecodeParameter("RANNodeName", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANNodeName{}, 0, err
	}
	return v, len(data), err
}
