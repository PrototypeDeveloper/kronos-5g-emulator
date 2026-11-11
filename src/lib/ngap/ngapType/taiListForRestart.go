// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAIListForRestart represents ASN.1 TAIListForRestart.
type TAIListForRestart struct {
	List []TAI `aper:"sizeLB:1,sizeUB:2048"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIListForRestart) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAIListForRestart", v)
}

// DecodeTAIListForRestart decodes a standalone APER value and reports the consumed octets.
func DecodeTAIListForRestart(data []byte) (TAIListForRestart, int, error) {
	var v TAIListForRestart
	err := ngapcodec.DecodeParameter("TAIListForRestart", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIListForRestart{}, 0, err
	}
	return v, len(data), err
}
