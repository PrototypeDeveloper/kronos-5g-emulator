// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// XnGTPTLAs represents ASN.1 XnGTP-TLAs.
type XnGTPTLAs struct {
	List []TransportLayerAddress `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v XnGTPTLAs) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("XnGTP-TLAs", v) }

// DecodeXnGTPTLAs decodes a standalone APER value and reports the consumed octets.
func DecodeXnGTPTLAs(data []byte) (XnGTPTLAs, int, error) {
	var v XnGTPTLAs
	err := ngapcodec.DecodeParameter("XnGTP-TLAs", data, &v)
	if ngapcodec.IsFatal(err) {
		return XnGTPTLAs{}, 0, err
	}
	return v, len(data), err
}
