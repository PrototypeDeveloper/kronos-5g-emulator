// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AMFName represents ASN.1 AMFName.
type AMFName struct {
	Value string `aper:"sizeExt,sizeLB:1,sizeUB:150"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFName) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("AMFName", v) }

// DecodeAMFName decodes a standalone APER value and reports the consumed octets.
func DecodeAMFName(data []byte) (AMFName, int, error) {
	var v AMFName
	err := ngapcodec.DecodeParameter("AMFName", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFName{}, 0, err
	}
	return v, len(data), err
}
