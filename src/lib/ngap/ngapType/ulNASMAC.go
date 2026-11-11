// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ULNASMAC represents ASN.1 UL-NAS-MAC.
type ULNASMAC struct {
	Value BitString `aper:"sizeLB:16,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ULNASMAC) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("UL-NAS-MAC", v) }

// DecodeULNASMAC decodes a standalone APER value and reports the consumed octets.
func DecodeULNASMAC(data []byte) (ULNASMAC, int, error) {
	var v ULNASMAC
	err := ngapcodec.DecodeParameter("UL-NAS-MAC", data, &v)
	if ngapcodec.IsFatal(err) {
		return ULNASMAC{}, 0, err
	}
	return v, len(data), err
}
