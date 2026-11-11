// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAIListForWarning represents ASN.1 TAIListForWarning.
type TAIListForWarning struct {
	List []TAI `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIListForWarning) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAIListForWarning", v)
}

// DecodeTAIListForWarning decodes a standalone APER value and reports the consumed octets.
func DecodeTAIListForWarning(data []byte) (TAIListForWarning, int, error) {
	var v TAIListForWarning
	err := ngapcodec.DecodeParameter("TAIListForWarning", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIListForWarning{}, 0, err
	}
	return v, len(data), err
}
