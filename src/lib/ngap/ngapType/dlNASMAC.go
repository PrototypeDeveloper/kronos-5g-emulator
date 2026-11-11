// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DLNASMAC represents ASN.1 DL-NAS-MAC.
type DLNASMAC struct {
	Value BitString `aper:"sizeLB:16,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DLNASMAC) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("DL-NAS-MAC", v) }

// DecodeDLNASMAC decodes a standalone APER value and reports the consumed octets.
func DecodeDLNASMAC(data []byte) (DLNASMAC, int, error) {
	var v DLNASMAC
	err := ngapcodec.DecodeParameter("DL-NAS-MAC", data, &v)
	if ngapcodec.IsFatal(err) {
		return DLNASMAC{}, 0, err
	}
	return v, len(data), err
}
