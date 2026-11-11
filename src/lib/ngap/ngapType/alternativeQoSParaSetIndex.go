// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AlternativeQoSParaSetIndex represents ASN.1 AlternativeQoSParaSetIndex.
type AlternativeQoSParaSetIndex struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AlternativeQoSParaSetIndex) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AlternativeQoSParaSetIndex", v)
}

// DecodeAlternativeQoSParaSetIndex decodes a standalone APER value and reports the consumed octets.
func DecodeAlternativeQoSParaSetIndex(data []byte) (AlternativeQoSParaSetIndex, int, error) {
	var v AlternativeQoSParaSetIndex
	err := ngapcodec.DecodeParameter("AlternativeQoSParaSetIndex", data, &v)
	if ngapcodec.IsFatal(err) {
		return AlternativeQoSParaSetIndex{}, 0, err
	}
	return v, len(data), err
}
