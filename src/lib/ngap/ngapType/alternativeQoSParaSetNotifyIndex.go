// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AlternativeQoSParaSetNotifyIndex represents ASN.1 AlternativeQoSParaSetNotifyIndex.
type AlternativeQoSParaSetNotifyIndex struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AlternativeQoSParaSetNotifyIndex) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AlternativeQoSParaSetNotifyIndex", v)
}

// DecodeAlternativeQoSParaSetNotifyIndex decodes a standalone APER value and reports the consumed octets.
func DecodeAlternativeQoSParaSetNotifyIndex(data []byte) (AlternativeQoSParaSetNotifyIndex, int, error) {
	var v AlternativeQoSParaSetNotifyIndex
	err := ngapcodec.DecodeParameter("AlternativeQoSParaSetNotifyIndex", data, &v)
	if ngapcodec.IsFatal(err) {
		return AlternativeQoSParaSetNotifyIndex{}, 0, err
	}
	return v, len(data), err
}
