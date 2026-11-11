// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CAGListforMDT represents ASN.1 CAGListforMDT.
type CAGListforMDT struct {
	List []CAGListforMDTItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CAGListforMDT) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("CAGListforMDT", v) }

// DecodeCAGListforMDT decodes a standalone APER value and reports the consumed octets.
func DecodeCAGListforMDT(data []byte) (CAGListforMDT, int, error) {
	var v CAGListforMDT
	err := ngapcodec.DecodeParameter("CAGListforMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return CAGListforMDT{}, 0, err
	}
	return v, len(data), err
}
