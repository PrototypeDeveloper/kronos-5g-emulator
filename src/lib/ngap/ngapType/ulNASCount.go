// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ULNASCount represents ASN.1 UL-NAS-Count.
type ULNASCount struct {
	Value BitString `aper:"sizeLB:5,sizeUB:5"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ULNASCount) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("UL-NAS-Count", v) }

// DecodeULNASCount decodes a standalone APER value and reports the consumed octets.
func DecodeULNASCount(data []byte) (ULNASCount, int, error) {
	var v ULNASCount
	err := ngapcodec.DecodeParameter("UL-NAS-Count", data, &v)
	if ngapcodec.IsFatal(err) {
		return ULNASCount{}, 0, err
	}
	return v, len(data), err
}
