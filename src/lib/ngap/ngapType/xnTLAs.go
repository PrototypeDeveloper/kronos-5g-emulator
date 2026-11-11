// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// XnTLAs represents ASN.1 XnTLAs.
type XnTLAs struct {
	List []TransportLayerAddress `aper:"sizeLB:1,sizeUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v XnTLAs) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("XnTLAs", v) }

// DecodeXnTLAs decodes a standalone APER value and reports the consumed octets.
func DecodeXnTLAs(data []byte) (XnTLAs, int, error) {
	var v XnTLAs
	err := ngapcodec.DecodeParameter("XnTLAs", data, &v)
	if ngapcodec.IsFatal(err) {
		return XnTLAs{}, 0, err
	}
	return v, len(data), err
}
