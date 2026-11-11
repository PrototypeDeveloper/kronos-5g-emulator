// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TACListInNRNTN represents ASN.1 TACListInNRNTN.
type TACListInNRNTN struct {
	List []TAC `aper:"sizeLB:1,sizeUB:12"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TACListInNRNTN) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TACListInNRNTN", v)
}

// DecodeTACListInNRNTN decodes a standalone APER value and reports the consumed octets.
func DecodeTACListInNRNTN(data []byte) (TACListInNRNTN, int, error) {
	var v TACListInNRNTN
	err := ngapcodec.DecodeParameter("TACListInNRNTN", data, &v)
	if ngapcodec.IsFatal(err) {
		return TACListInNRNTN{}, 0, err
	}
	return v, len(data), err
}
