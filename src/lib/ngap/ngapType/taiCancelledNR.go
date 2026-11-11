// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAICancelledNR represents ASN.1 TAICancelledNR.
type TAICancelledNR struct {
	List []TAICancelledNRItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAICancelledNR) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAICancelledNR", v)
}

// DecodeTAICancelledNR decodes a standalone APER value and reports the consumed octets.
func DecodeTAICancelledNR(data []byte) (TAICancelledNR, int, error) {
	var v TAICancelledNR
	err := ngapcodec.DecodeParameter("TAICancelledNR", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAICancelledNR{}, 0, err
	}
	return v, len(data), err
}
