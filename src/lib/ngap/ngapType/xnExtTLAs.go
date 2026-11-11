// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// XnExtTLAs represents ASN.1 XnExtTLAs.
type XnExtTLAs struct {
	List []XnExtTLAItem `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v XnExtTLAs) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("XnExtTLAs", v) }

// DecodeXnExtTLAs decodes a standalone APER value and reports the consumed octets.
func DecodeXnExtTLAs(data []byte) (XnExtTLAs, int, error) {
	var v XnExtTLAs
	err := ngapcodec.DecodeParameter("XnExtTLAs", data, &v)
	if ngapcodec.IsFatal(err) {
		return XnExtTLAs{}, 0, err
	}
	return v, len(data), err
}
